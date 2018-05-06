package main

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

// ConvertStringValue convert string value to proper json type
func ConvertStringValue(stringValue string, jsonType string) (interface{}, error) {
	if jsonType == "number" {
		return strconv.Atoi(stringValue)
	} else if jsonType == "boolean" {
		return strconv.ParseBool(stringValue)
	} else {
		return stringValue, nil
	}
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
		value, err := ConvertStringValue(lineValues[col], header.Type)
		if err == nil {
			jsonObj[header.Name] = value
		} else {
			jsonObj[header.Name] = header.Default
		}
	}

	return jsonObj
}
