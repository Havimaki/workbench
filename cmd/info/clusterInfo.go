package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ====== helpers
func kubectl(dummy string) (string, error) {
	return dummy, nil
}

// ====== handlers
var clusterInfoCmd = &cobra.Command{
	Use:   "clusterInfo",
	Short: "Cluster info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clusterInfo called")

		if resp, err := kubectl("dummy_string"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	InfoCmd.AddCommand(clusterInfoCmd)
}
