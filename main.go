package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// CounterStruct Holds the number of occurrences for imports and template calls
type CounterStruct struct {
	impt, template int
}

var (
	flags      []cli.Flag
	dir        string
	components []string
	output     bool
	cnLength   int
)

func run() {
	files, err := FilesWalk(Concat(dir, "/src"), "*.vue")
	var componentNames []string
	cnLength = 0
	componentMap := make(map[string]*CounterStruct)

	// init component map
	for _, filePath := range files {
		c := GetComponentName(filePath)

		if strings.ToLower(c) == "app.vue" {
			continue
		}

		componentMap[c] = &CounterStruct{0, 0}
	}

	if err != nil {
		fmt.Println("ERROR")
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

					// if templateOccurrence > 0 {
					// 	fmt.Printf("Found %s %d times in %s\n", templateName, templateOccurrence, GetComponentName(filePath))
					// }

					componentMap[name].template += templateOccurrence
				}
			}
		}

		PrintResults(componentMap)
	}
}

func info(app *cli.App) {
	app.Name = "Housekeeper 🧹 "
	app.Usage = "Keep track how often your Vue components are used."
	app.Author = "Bored Chinese"
	app.Version = "1.0"
}

func commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run housekeeping for all components.",
			Action: func(c *cli.Context) {
				fmt.Println(dir)
				run()
				fmt.Printf("✨ Done. Checked %d file(s).\n", cnLength)
			},
		},
	}
}

func init() {
	flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Usage:       "Set the directory of your Vue project.",
			Destination: &dir,
		},
		// cli.BoolFlag{
		// 	Name:        "output, o",
		// 	Usage:       "Output the data to data.csv.",
		// 	Destination: &output,
		// },
		// cli.StringSliceFlag{
		// 	Name:        "components, c",
		// 	Usage:       "List of component names (without file extension) that you want to specifically find.",
		// 	Destination: &components,
		// },
	}
}

func main() {
	app := cli.NewApp()
	app.Flags = flags
	info(app)
	commands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
