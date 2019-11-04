package main

import (
	"fmt"
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
		Logger(fmt.Sprintf("Component: %s - File: %s", name, fileName))
		c[fileName] = &ComponentStruct{0, 0, name, filePath}
	}
}

// Analyzer looks through each line to find import statements and template calls
func Analyzer(filePath string, c map[string]*ComponentStruct) {
	data, _ := ReadFile(filePath)

	tStatus, tData := GetTemplateData(data)
	iStatus, iData := GetScriptData(data)

	Logger(tStatus)
	Logger(iStatus)
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

	Logger("Initializing componentMap")
	initComponentMap(files, componentMap)
	Logger("Done.")

	if err != nil {
		Logger("Error in FilesWalk")
		panic(err)
	} else {
		for _, filePath := range files {
			Logger(Concat("Reading ", filePath))
			Analyzer(filePath, componentMap)
			Logger(Concat("Successfully read ", filePath))
		}
	}

	PrintResults(componentMap)
}
