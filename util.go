package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/gookit/color.v1"
)

// FilesWalk recursively walks through files finding ones by extension
func FilesWalk(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// PrintStringArray prints all strings in the array
func PrintStringArray(files []string) {
	for _, str := range files {
		fmt.Println(str)
	}
}

// ReadFile reads a file! :)
func ReadFile(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	return string(f), err
}

// GetComponentName fetches the component name of a given file path
func GetComponentName(filePath string) string {
	s := strings.Split(filePath, "/")
	return s[len(s)-1]
}

// Concat merges two strings together
func Concat(a, b string) string {
	var str strings.Builder

	str.WriteString(a)
	str.WriteString(b)
	return str.String()
}

// RemoveExtension removes any file extensions from a name
func RemoveExtension(name string) string {
	s := strings.Split(name, ".")
	return s[0]
}

// PrintResults prints the component counter results
func PrintResults(components map[string]*CounterStruct) {
	for k, v := range components {
		k = RemoveExtension(k)
		fmt.Print(k, " ")

		if v.impt > 0 {
			color.Green.Printf("%d import(s)", v.impt)
		} else if v.impt == 0 {
			color.Red.Printf("%d import(s)", v.impt)
		}

		fmt.Print(" - ")

		if v.template > 0 {
			color.Green.Printf("%d call(s)", v.template)
		} else if v.template == 0 {
			color.Red.Printf("%d call(s)", v.template)
		}

		fmt.Println()
	}
}
