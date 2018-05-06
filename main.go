package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// args := os.Args[1:]

	headerArg := flag.String("header", "", "List of headers key")
	inputFile := flag.String("input", "", "Input csv file to convert")
	flag.Parse()

	headerValues := strings.TrimSpace(*headerArg)
	headers := []Header{}
	if headerValues != "" {
		headerKeys := strings.Split(headerValues, ",")
		for _, key := range headerKeys {
			headerParts := strings.SplitN(key, ":", 3)
			headerName, headerType, defaultValueInString := headerParts[0], headerParts[1], headerParts[2]
			defaultValue, _ := ConvertStringValue(defaultValueInString, headerType)
			header := Header{Name: headerName, Type: headerType, Default: defaultValue}
			headers = append(headers, header)
		}
	}

	var bytes []byte
	var err error
	if *inputFile != "" {
		bytes, err = ioutil.ReadFile(*inputFile)
	} else {
		bytes, err = ioutil.ReadAll(os.Stdin)
	}

	if err == nil {
		fmt.Println(Convert(string(bytes), headers))
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
