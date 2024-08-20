package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command, this rereads the configuration file and updates the running
// process compose server with the new configuration
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Reread the configuration file and update the running process compose server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("src/cmd/refresh.go \n")
		fmt.Printf("%+v\n", getClient().GetProjectConfig())
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
