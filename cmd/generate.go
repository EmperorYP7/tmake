/*
Package cmd tmake
Copyright © 2021 Yash Pandey

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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:        "generate",
	Aliases:    []string{"gen"},
	SuggestFor: []string{"genrate", "make", "create"},
	Short:      "Lets you generate templates",
	Long: `
The command lets you generate templates based on remote or local
sources if they're valid otherwise, will throw an error.
`,
	Example:   "tmake generate [arguments] -flags=value",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires at least two arguments")
		}
		return nil
	},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            map[string]string{},
	Version:                "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running generate command")
		fmt.Println(cmd.Flags().GetString("dest"))
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
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	var destination = ""
	generateCmd.Flags().StringVarP(&destination ,"dest", "d", "", "Destination of generated templates")
}
