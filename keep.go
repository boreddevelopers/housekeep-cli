package main

import (
	"strings"
)

// ComponentStruct Holds the number of occurrences for imports and template calls
type ComponentStruct struct {
	impt, template int
}

// Keep contains all the logic for browsing through .vue files.
func Keep() {
	files, err := FilesWalk(dir, "*.vue")
	var componentNames []string
	cnLength = 0
	componentMap := make(map[string]*ComponentStruct)

	// init component map
	for _, filePath := range files {
		c := GetComponentName(filePath)

		if strings.ToLower(c) == "app.vue" {
			continue
		}

		componentMap[c] = &ComponentStruct{0, 0}
	}

	if err != nil {
		panic("Unable to open file.")
	} else {
		for _, filePath := range files {
			componentName := GetComponentName(filePath)
			// Ignore App.vue because it's common in every project.
			if strings.ToLower(componentName) == "app.vue" {
				continue
			}

			componentNames = append(componentNames, componentName)
			cnLength = len(componentNames)
		}

		for _, filePath := range files {
			data, err := ReadFile(filePath)

			if err == nil {
				for _, name := range componentNames {
					keyword := RemoveExtension(name)
					importName := Concat("import ", keyword)
					if strings.Contains(data, importName) {
						componentMap[name].impt++
					}
					templateName := Concat("<", keyword)
					templateOccurrence := strings.Count(data, templateName)
					componentMap[name].template += templateOccurrence
				}
			}
		}

		if toPrint {
			PrintResults(componentMap)
		}
	}
}
