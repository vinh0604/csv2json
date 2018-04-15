package csv2json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name     interface{} `json:"name"`
	Age      interface{} `json:"age"`
	Location interface{} `json:"location"`
	Active   interface{} `json:"active"`
}

func TestConvertShouldUseFirstRowAsHeadersIfHeadersAreEmpty(t *testing.T) {
	assert := assert.New(t)

	csv := `name,age,location
John Doe,21,Singapore
James Smith,25,England
Joe Biden,75,USA`
	result := Convert(csv, []Header{})

	actual := []Person{}
	expected := []Person{}
	json.Unmarshal([]byte(`[
		{"name": "John Doe", "age": "21", "location": "Singapore"},
		{"name": "James Smith", "age": "25", "location": "England"},
		{"name": "Joe Biden", "age": "75", "location": "USA"}
	]`), &expected)
	json.Unmarshal([]byte(result), &actual)
	assert.Equal(expected, actual)
}

func TestConvertShouldIgnoreEmptyLines(t *testing.T) {
	assert := assert.New(t)

	csv := `name,age,location
John Doe,21,Singapore

Joe Biden,75,USA
`
	result := Convert(csv, []Header{})

	actual := []Person{}
	expected := []Person{}
	json.Unmarshal([]byte(`[
		{"name": "John Doe", "age": "21", "location": "Singapore"},
		{"name": "Joe Biden", "age": "75", "location": "USA"}
	]`), &expected)
	json.Unmarshal([]byte(result), &actual)
	assert.Equal(expected, actual)
}

func TestConvertShouldUseHeaderMetaDataIfProvided(t *testing.T) {
	assert := assert.New(t)

	csv := `John Doe,21,Singapore
Joe Biden,75,USA`
	headers := []Header{
		Header{Name: "name"},
		Header{Name: "age"},
		Header{Name: "location"},
	}
	result := Convert(csv, headers)

	actual := []Person{}
	expected := []Person{}
	json.Unmarshal([]byte(`[
		{"name": "John Doe", "age": "21", "location": "Singapore"},
		{"name": "Joe Biden", "age": "75", "location": "USA"}
	]`), &expected)
	json.Unmarshal([]byte(result), &actual)
	assert.Equal(expected, actual)
}

func TestConvertShouldParseWithDataTypeSpecifiedInHeader(t *testing.T) {
	assert := assert.New(t)

	csv := `John Doe,21,true
Joe Biden,75,false`
	headers := []Header{
		Header{Name: "name"},
		Header{Name: "age", Type: "number"},
		Header{Name: "active", Type: "boolean"},
	}
	result := Convert(csv, headers)

	actual := []Person{}
	expected := []Person{}
	json.Unmarshal([]byte(`[
		{"name": "John Doe", "age": 21, "active": true},
		{"name": "Joe Biden", "age": 75, "active": false}
	]`), &expected)
	json.Unmarshal([]byte(result), &actual)
	assert.Equal(expected, actual)
}

func TestConvertShouldFallBackToDefaultValuesSpecifiedInHeaderIfProvidedValueIsInvalid(t *testing.T) {
	assert := assert.New(t)

	csv := `John Doe,,true
Joe Biden,75,`
	headers := []Header{
		Header{Name: "name"},
		Header{Name: "age", Type: "number"},
		Header{Name: "active", Type: "boolean", Default: true},
	}
	result := Convert(csv, headers)

	actual := []Person{}
	expected := []Person{}
	json.Unmarshal([]byte(`[
		{"name": "John Doe", "age": null, "active": true},
		{"name": "Joe Biden", "age": 75, "active": true}
	]`), &expected)
	json.Unmarshal([]byte(result), &actual)
	assert.Equal(expected, actual)
}
