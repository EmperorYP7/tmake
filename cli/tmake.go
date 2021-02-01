package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "Template Maker"
  app.HelpName = "tmake"
  app.Usage = "A Command Line Interface for generating templates"
  app.Commands = []cli.Command{
    cli.Command{
    	Name:        "generate",
    	ShortName:   "gen",
    	Aliases:     []string{"make"},
    	Usage:       "Generates file/template in the provided folder",
    	UsageText:   "generate `FILEPATH`",
    	Description: "Generates file/template in the provided folder",
    	ArgsUsage:   "[Arrgh]",
    	Category:    "Generate",
    	BashComplete: func(c *cli.Context) {
      },
      Before: func(c *cli.Context) error {
        return nil
      },
      After: func(c *cli.Context) error {
        return nil
      },
    	Subcommands:            []cli.Command{},
    	Flags:                  []cli.Flag{},
    	SkipFlagParsing:        false,
    	SkipArgReorder:         false,
    	HideHelp:               false,
    	Hidden:                 false,
    	UseShortOptionHandling: false,
    	HelpName:               "generate",
    	CustomHelpTemplate:     "",
    },
    cli.Command{
    	Name:        "remote",
    	ShortName:   "rt",
    	Aliases:     []string{"ext"},
    	Usage:       "Showing all the templates imported from external sources",
    	UsageText:   "",
    	Description: "",
    	ArgsUsage:   "",
    	Category:    "Remote",
    	BashComplete: func( *cli.Context) {
    	},
    	Before: func( *cli.Context) error {
        return nil
      },
    	After: func( *cli.Context) error {
        return nil
      },
    	Action: nil,
    	Subcommands:            []cli.Command{},
    	Flags:                  []cli.Flag{},
    	SkipFlagParsing:        false,
    	SkipArgReorder:         false,
    	HideHelp:               false,
    	Hidden:                 false,
    	UseShortOptionHandling: false,
    	HelpName:               "",
    	CustomHelpTemplate:     "",
    },
  }
  app.Run(os.Args)
}
