package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var build = "1" // build number set at compile time

func main() {
	app := &cli.App{
		Name:  "Drone-Sonar-Plugin",
		Usage: "Drone plugin to integrate with SonarQube.",
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("key")) // Print the "key" value as an example
			return nil
		},
		Version: "1.0." + build,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "key",
				Usage:   "project key",
				EnvVars: []string{"DRONE_REPO"},
			},
			&cli.StringFlag{
				Name:    "name",
				Usage:   "project name",
				EnvVars: []string{"DRONE_REPO"},
			},
			&cli.StringFlag{
				Name:    "host",
				Usage:   "SonarQube host",
				EnvVars: []string{"PLUGIN_SONAR_HOST"},
			},
			&cli.StringFlag{
				Name:    "token",
				Usage:   "SonarQube token",
				EnvVars: []string{"PLUGIN_SONAR_TOKEN"},
			},
			// Advanced parameters
			&cli.StringFlag{
				Name:    "ver",
				Usage:   "Project version",
				EnvVars: []string{"DRONE_BUILD_NUMBER"},
			},
			&cli.StringFlag{
				Name:    "branch",
				Usage:   "Project branch",
				EnvVars: []string{"DRONE_BRANCH"},
			},
			&cli.StringFlag{
				Name:    "timeout",
				Usage:   "Web request timeout",
				Value:   "60",
				EnvVars: []string{"PLUGIN_TIMEOUT"},
			},
			&cli.StringFlag{
				Name:    "sources",
				Usage:   "analysis sources",
				Value:   ".",
				EnvVars: []string{"PLUGIN_SOURCES"},
			},
			&cli.StringFlag{
				Name:    "inclusions",
				Usage:   "code inclusions",
				EnvVars: []string{"PLUGIN_INCLUSIONS"},
			},
			&cli.StringFlag{
				Name:    "exclusions",
				Usage:   "code exclusions",
				EnvVars: []string{"PLUGIN_EXCLUSIONS"},
			},
			&cli.StringFlag{
				Name:    "level",
				Usage:   "log level",
				Value:   "INFO",
				EnvVars: []string{"PLUGIN_LEVEL"},
			},
			&cli.BoolFlag{
				Name:    "showProfiling",
				Usage:   "showProfiling during analysis",
				Value:   false,
				EnvVars: []string{"PLUGIN_SHOWPROFILING"},
			},
			&cli.BoolFlag{
				Name:    "branchAnalysis",
				Usage:   "execute branchAnalysis",
				EnvVars: []string{"PLUGIN_BRANCHANALYSIS"},
			},
			&cli.BoolFlag{
				Name:    "usingProperties",
				Usage:   "using sonar-project.properties",
				EnvVars: []string{"PLUGIN_USINGPROPERTIES"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
