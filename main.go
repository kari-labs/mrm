package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mUltI rEpO ManaGER"
	app.Usage = "manage repos"

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a repo to the config and clone it",
			Action:  addRepo,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Usage: "Directory to save the repo in",
				},
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove a repo from the config and delete it",
			Action:  removeRepo,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "keep, k",
					Usage: "Remove the repo from the config but not from the filesystem",
				},
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Force delete repo regardless of unpushed or uncommited changes",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
