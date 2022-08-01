/*
Copyright © 2022 Shady Aziza

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

var (
	fstRegex = ConvetionalCommitsRegex{value: `\A(((Initial commit)|(Merge [^\r\n]+)|((build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test)(\(\w+\))?!?: [^\r\n]+((\r|\n|\r\n)((\r|\n|\r\n)[^\r\n]+)+)*))(\r|\n|\r\n)?)\z`, ID: 1}
)

type ConvetionalCommitsRegex struct {
	value string
	ID    int
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govern",
	Short: "Enforce conventional commits",
	Long: `Govern enforces the use of conventional commits in a team project:

A CLI tool for enforcing conventional commits https://www.conventionalcommits.org/en/v1.0.0/,
and can be easily integrated with your git hooks`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")

		validateMessage(message)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.govern.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("message", "", "a commit message")
}

func validateMessage(message string) {

	r, err := regexp.MatchString(fstRegex.value, message)
	if err != nil {
		fmt.Printf("An error has occured with CLI tool.")
	}
	if r {

		os.Exit(0)
	}
	fmt.Printf("BUSTED.\n")
	os.Exit(1)

}
