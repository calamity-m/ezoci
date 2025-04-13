package search

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/ezoci"
	"github.com/calamity-m/ezoci/pkg/search"
	"github.com/spf13/cobra"
)

var UrlOverride string

func NewSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "search",
		Long:  `search`,
	}

	cmd.PersistentFlags().StringVarP(&UrlOverride, "url", "u", "", "override registry url")
	cmd.AddCommand(newGhcrCommand())
	cmd.AddCommand(newDockerCommand())
	cmd.AddCommand(newGcrCommand())
	cmd.AddCommand(newGitlabCommand())
	cmd.AddCommand(newHarborCommand())

	return cmd
}

func newGhcrCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ghcr",
		Short: "ghcr",
		Long:  `ghcr`,
		Args:  cobra.ExactArgs(1),
		Run: func(ccmd *cobra.Command, args []string) {
			out, err := ezoci.Search(
				&search.SearchOpts{
					Provider:    search.Ghcr,
					UrlOverride: UrlOverride,
					Path:        args[0],
				},
			)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to search GHCR registry - %s", err), slog.Any("error", err))
				os.Exit(1)
			}

			slog.Info(out)
		},
	}
}

func newDockerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "docker",
		Short: "docker",
		Long:  `docker`,
		Args:  cobra.ExactArgs(1),
		Run: func(ccmd *cobra.Command, args []string) {
			out, err := ezoci.Search(
				&search.SearchOpts{
					Provider:    search.Docker,
					UrlOverride: UrlOverride,
					Path:        args[0],
				},
			)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to search Docker registry - %s", err), slog.Any("error", err))
				os.Exit(1)
			}

			slog.Info(out)
		},
	}
}

func newGcrCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gcr",
		Short: "gcr",
		Long:  `gcr`,
		Args:  cobra.ExactArgs(1),
		Run: func(ccmd *cobra.Command, args []string) {
			out, err := ezoci.Search(
				&search.SearchOpts{
					Provider:    search.Gcr,
					UrlOverride: UrlOverride,
					Path:        args[0],
				},
			)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to search GCR registry - %s", err), slog.Any("error", err))
				os.Exit(1)
			}

			slog.Info(out)
		},
	}
}

func newGitlabCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gitlab",
		Short: "gitlab",
		Long:  `gitlab`,
		Args:  cobra.ExactArgs(1),
		Run: func(ccmd *cobra.Command, args []string) {
			out, err := ezoci.Search(
				&search.SearchOpts{
					Provider:    search.Gitlab,
					UrlOverride: UrlOverride,
					Path:        args[0],
				},
			)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to search Gitlab registry - %s", err), slog.Any("error", err))
				os.Exit(1)
			}

			slog.Info(out)
		},
	}
}

func newHarborCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "harbor",
		Short: "harbor",
		Long:  `harbor`,
		Args:  cobra.ExactArgs(1),
		Run: func(ccmd *cobra.Command, args []string) {
			out, err := ezoci.Search(
				&search.SearchOpts{
					Provider:    search.Harbor,
					UrlOverride: UrlOverride,
					Path:        args[0],
				},
			)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to search Harbor registry - %s", err), slog.Any("error", err))
				os.Exit(1)
			}

			slog.Info(out)
		},
	}
}
