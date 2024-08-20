package cmd

import (
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command, this rereads the configuration file and updates the running
// process compose server with the new configuration
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Reread the configuration file and update the running process compose server",
	Run: func(cmd *cobra.Command, args []string) {
		println("refresh called")
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

}
