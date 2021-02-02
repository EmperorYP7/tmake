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
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:        "generate",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "Lets you generate templates",
	Long: `
The command lets you generate templates based on remote or local
sources if they're valid otherwise, will throw an error.
`,
	Example:   "tmake generate [sub-commands] -flags=value",
	ValidArgs: []string{},
	// Args: func(cmd *cobra.Command, args []string) error {
	// },
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            map[string]string{},
	Version:                "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
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
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
