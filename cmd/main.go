package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/ezoci/pkg/cmd"
	cmd_search "github.com/calamity-m/ezoci/pkg/cmd/search"
	"github.com/calamity-m/ezoci/pkg/logging"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ezoci",
	Short: "ezoci",
	Long:  "ezoci",
	PersistentPreRun: func(ccmd *cobra.Command, args []string) {
		logger := slog.New(
			logging.NewInterceptHandler(
				os.Stderr,
				nil,
				cmd.PrettyPrint,
				cmd.Verbose,
			),
		)

		slog.SetDefault(logger)
	},
}

func Execute() error {

	rootCmd.AddCommand(cmd_search.NewSearchCommand())

	rootCmd.PersistentFlags().BoolVarP(&cmd.Verbose, "verbose", "v", false, "Verbose")
	rootCmd.PersistentFlags().BoolVarP(&cmd.PrettyPrint, "pretty-print", "p", false, "Prettified output")

	return rootCmd.Execute()
}

func main() {
	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
