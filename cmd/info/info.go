package info

import (
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "All things information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
