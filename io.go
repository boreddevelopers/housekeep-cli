package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadFile reads a file! :)
func ReadFile(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	return string(f), err
}

// ReadAndSplitLines does stuff
func ReadAndSplitLines(filePath string, toPrint bool) []string {
	data, _ := ReadFile(filePath)
	arrData := strings.Split(data, "\n")

	if toPrint {
		for _, d := range arrData {
			fmt.Println(d)
		}
	}

	return arrData
}
