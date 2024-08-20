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
		// This is a new runner with the latest updates from the config:
		// runner := getProjectRunner(args, *pcFlags.NoDependencies, "", []string{})
		// fmt.Printf("NEW RUNNER: %+v\n", runner)

		// I still need to somehow get the running project:
		projects := getClient().GetProjectConfig();
		fmt.Printf("REFRESH RESULT: %+v\n", projects)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
