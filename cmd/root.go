package cmd

import (
	"github.com/spf13/cobra"
)

var debug bool
var pretty bool
var structured bool
var rootCmd = &cobra.Command{
	Use:   "ezoci",
	Short: "ezoci",
	Long:  "ezoci",
}

func Execute() error {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "v", false, "Debug print")
	rootCmd.PersistentFlags().BoolVarP(&pretty, "pretty-print", "p", false, "prettify output with colours and log level")
	rootCmd.PersistentFlags().BoolVarP(&structured, "structured-print", "S", false, "Display logging in a structured format")

	// Register sub-commands
	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(parentCmd)

	return rootCmd.Execute()
}
