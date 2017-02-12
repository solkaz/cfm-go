package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cfm"
	app.Usage = "config file manager written in go"

	app.Commands = []cli.Command{
		{
			Name: "list",
			Aliases: []string{"ls"},
			Usage: "list all (or one) aliases and their mapped filepaths",
		},
	}

	app.Run(os.Args)
}
