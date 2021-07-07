/*
Package cmd tmake
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:        "remote",
	Aliases:    []string{"rmt"},
	SuggestFor: []string{},
	Short:      "Shows the list of all remote templates",
	Long: `
The command is associated with all the tasks related to generation, deletion,
editing and copying of templates from remote sources. By default it shows
the list of all remote templates.
`,
	Example:   "",
	ValidArgs: []string{},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            map[string]string{},
	Run: func(cmd *cobra.Command, args []string) {
		fileTextLines, err := ReadRemoteFile("./store/remote.csv")
		if err != nil {
			fmt.Print(err)
		}
		for _, eachline := range fileTextLines {
			fmt.Println(eachline)
		}
	},
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
