package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func check(e error) {
	if e != nil {
		fmt.Println("ERROR: Unable to open file.")
	}
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

	// if err != nil {
	// 	fmt.Println("[ERROR] Unable to open file")
	// } else {
	// 	fmt.Println(string(f))
	// }

	return string(f), err
}

func getComponentName(filePath string) string {
	s := strings.Split(filePath, "/")
	return s[len(s)-1]
}

func concat(a, b string) string {
	var str strings.Builder

	str.WriteString(a)
	str.WriteString(b)
	return str.String()
}

func removeExtension(name string) string {
	s := strings.Split(name, ".")
	return s[0]
}

func main() {
	files, err := FilesWalk("../bposeats.com/src/", "*.vue")
	var componentNames []string
	var componentMap = make(map[string]int)

	// init component map
	for _, filePath := range files {
		c := getComponentName(filePath)

		if strings.ToLower(c) == "app.vue" {
			continue
		}

		componentMap[c] = 0
	}

	if err != nil {
		fmt.Println("ERROR")
	} else {
		for _, filePath := range files {
			componentName := getComponentName(filePath)
			// Ignore App.vue because it's common in every project.
			if strings.ToLower(componentName) == "app.vue" {
				continue
			}

			componentNames = append(componentNames, componentName)
		}

		// TODO: Find words in file
		for _, name := range componentNames {
			// fmt.Println(name)
			// 2. import statement
			for _, filePath := range files {
				data, err := ReadFile(filePath)
				if err == nil {
					keyword := removeExtension(name)
					if strings.Contains(data, keyword) {
						componentMap[name]++
					}
				}
			}
			fmt.Printf("%s: %d\n", removeExtension(name), componentMap[name])
		}
	}
}
