package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	flags      []cli.Flag
	dir        string
	components []string
	output     bool
	cnLength   int
	toPrint    bool
	toLog      bool
)

func info(app *cli.App) {
	app.Name = "Housekeep 🧹 "
	app.Usage = "Keep track how often your Vue components are used."
	app.Author = "Bored Chinese"
	app.Version = "1.0.1"
}

func commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run housekeeping for all components.",
			Action: func(c *cli.Context) {
				Keep()
				fmt.Print("\n✨ Housekeeping finished.\n")
			},
		},
	}
}

func init() {
	flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Usage:       "Set the directory of your Vue project.",
			Value:       GetCWD(),
			Destination: &dir,
		},
		cli.BoolFlag{
			Name:        "print, p",
			Usage:       "Print the tallies in CLI",
			Destination: &toPrint,
		},
		cli.BoolFlag{
			Name:        "log, l",
			Usage:       "Log debugging information if something isn't working.",
			Destination: &toLog,
		},
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
