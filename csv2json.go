package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"strconv"
	"strings"
)

// Header csv header metadata
type Header struct {
	Name, Type string
	Default    interface{}
}

// Convert convert csv to json with optional headers
func Convert(data string, headers []Header) string {
	reader := csv.NewReader(strings.NewReader(data))
	lines, _ := reader.ReadAll()

	result := []map[string]interface{}{}
	for rowIndex, line := range lines {
		if len(headers) == 0 && rowIndex == 0 {
			headers = convertFirstLineToHeader(line)
			continue
		}

		result = append(result, convertLineToJSONObject(line, headers))
	}

	jsonBytes, _ := json.Marshal(result)
	return string(jsonBytes[:])
}

func convertFirstLineToHeader(lineValues []string) []Header {
	headers := []Header{}
	for _, key := range lineValues {
		headers = append(headers, Header{Name: key})
	}
	return headers
}

func convertLineToJSONObject(lineValues []string, headers []Header) map[string]interface{} {
	jsonObj := make(map[string]interface{})
	for col, header := range headers {
		var value interface{}
		var err error
		if header.Type == "number" {
			value, err = strconv.Atoi(lineValues[col])
		} else if header.Type == "boolean" {
			value, err = strconv.ParseBool(lineValues[col])
		} else {
			value, err = lineValues[col], nil
		}

		if err == nil {
			jsonObj[header.Name] = value
		} else {
			jsonObj[header.Name] = header.Default
		}
	}

	return jsonObj
}
