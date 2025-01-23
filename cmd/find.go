package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/ezoci/pkg/logging"
	"github.com/spf13/cobra"
)

var (
	findCmd = &cobra.Command{
		Use:   "fd",
		Short: "find",
		Long:  `find`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("find")

			// Initialize logging output
			var logger *slog.Logger
			if structured {
				logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
			} else {
				logger = slog.New(logging.NewInterceptedTextHandler(os.Stderr, pretty, debug, nil))
			}

			slog.SetDefault(logger)
			slog.Info(";)")
			slog.Info(fmt.Sprintf("%v", structured))
			slog.Debug("debug")
			slog.Error("err")
			slog.Info("info")
		},
	}
)
