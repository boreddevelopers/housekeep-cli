package main

import (
	"strings"
)

// ComponentStruct Holds the number of occurrences for imports and template calls
type ComponentStruct struct {
	impt     int    // Counter for the number of times it is called in import statements
	template int    // Counter for the number of times it is called in a template
	name     string // Component name
	filePath string // File path of the component including its file name & type
}

func initComponentMap(files []string, c map[string]*ComponentStruct) {
	for _, filePath := range files {
		name := GetComponentName(filePath)
		fileName := GetFileName(filePath)
		c[fileName] = &ComponentStruct{0, 0, name, filePath}
	}
}

// Analyzer looks through each line to find import statements and template calls
func Analyzer(filePath string, c map[string]*ComponentStruct) {
	data, _ := ReadFile(filePath)

	tData := GetTemplateData(data)
	iData := GetScriptData(data)

	for k, cStruct := range c {
		// Check template data
		tCount := strings.Count(tData, Concat("<", cStruct.name))
		c[k].template += tCount

		// Check import data
		iCount := strings.Count(iData, cStruct.name)
		c[k].impt += iCount
	}

}

// Keep contains all the logic for browsing through .vue files.
func Keep() {
	// Get all files with vue extension
	files, err := FilesWalk(dir, "*.vue")
	componentMap := make(map[string]*ComponentStruct)
	cnLength = len(files)

	initComponentMap(files, componentMap)

	if err != nil {
		panic(err)
	} else {
		for _, filePath := range files {
			Analyzer(filePath, componentMap)
		}
	}

	PrintResults(componentMap)
}
