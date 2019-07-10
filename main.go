package main

import (
	"log"
	"os"

	run "github.com/iwilltry42/k3d-tools/cli"
	"github.com/iwilltry42/k3d-tools/version"
	"github.com/urfave/cli"
)

// main represents the CLI application
func main() {

	// App Details
	app := cli.NewApp()
	app.Name = "k3d-tools"
	app.Usage = "Tools to help running k3d successfully!"
	app.Version = version.GetVersion()

	// commands that you can execute
	app.Commands = []cli.Command{
		{
			// check-tools verifies that docker is up and running
			Name:    "save-image",
			Aliases: []string{"save"},
			Usage:   "Check if docker is running",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "destination, dest, d",
					Value: "/images",
					Usage: "destination tar-file (optional)",
				},
				cli.StringFlag{
					Name:  "cluster, c",
					Value: "k3s-default",
					Usage: "name of the k3d cluster",
				},
			},
			Action: run.ImageSave,
		},
	}

	// run the whole thing
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
