package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ====== helpers
func diskUsage(directory string) (int, error) {
	usage := du.NewDiskUsage(directory)
	return fmt.Printf("Free disk space: %d in directory %s\n", usage.Free(), directory)
}

// ====== handlers
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Prints disk usage of the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		defaultDirectory := "."

		if dir := viper.GetViper().GetString("cmd.info.diskUsage.directory"); dir != "" {
			fmt.Println("Using config value")
			defaultDirectory = dir
		}
		diskUsage(defaultDirectory)
	},
}

func init() {
	// Add commands to Parent InfoCmd
	InfoCmd.AddCommand(diskUsageCmd)
}
