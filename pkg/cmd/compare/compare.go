package compare

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/ezoci"
	"github.com/spf13/cobra"
)

var Source string
var Others []string

func NewCompareCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "Compare a source container image against one or more other images",
		Run: func(cmd *cobra.Command, args []string) {
			slog.Debug("Starting Comparison")
			_, err := ezoci.Compare(Source, Others...)

			if err != nil {
				slog.Error(fmt.Sprintf("Failed to compare images - %s", err), slog.Any("error", err))
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(
		&Source,
		"source",
		"",
		"Source image to base comparison off",
	)
	cmd.Flags().StringSliceVar(
		&Others,
		"others",
		[]string{},
		"One or more other images to compare against the source",
	)
	cmd.MarkFlagRequired("source")
	cmd.MarkFlagRequired("others")

	return cmd
}
