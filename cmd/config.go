package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "All config-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetViper().GetString("cmd.info.diskUsage.DefaultDirectory"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
