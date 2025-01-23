package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parentCmd = &cobra.Command{
		Use:   "parent",
		Short: "verify an oci artifact's parent",
		Long:  `verify if an oci artifact is a parent of some other artifact`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("parent")
		},
	}
)
