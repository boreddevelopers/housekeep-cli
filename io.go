package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// DoesFileExist checks if the file exists
func DoesFileExist(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// AppendToFile appends data to an already existing file.
func AppendToFile(data, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(f, data)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// CreateNewFileWithData Create a new file and add data to it.
func CreateNewFileWithData(data, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Fprintln(f, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

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
