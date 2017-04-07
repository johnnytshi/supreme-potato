// pulls the latest Specs repo and parse the podspec.json files
package main

import (
	"fmt"
	"os"

	"github.com/supreme-potato/pod"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	(&cli.App{Commands: []*cli.Command{
		{
			Name:    "pod",
			Aliases: []string{"p"},
			Usage:   "let's make pod better",
			Subcommands: []*cli.Command{
				{
					Name:  "run",
					Usage: "process spec repo",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "repo-path",
							Value: "/Users/tingshi/.cocoapods/repos/master/master",
							Usage: "location of the spec repo",
						},
					},
					Action: func(c *cli.Context) error {
						pod.Run(c.String("repo-path"))
						return nil
					},
				},
				{
					Name:  "local",
					Usage: "setup local db tables",
					Action: func(c *cli.Context) error {
						pod.SetupLocal()
						return nil
					},
				},
			},
		},
	},
	}).Run(os.Args)

	fmt.Println("\n\n### be a rebel http://www.wikihow.com/Be-a-Rebel")
}
