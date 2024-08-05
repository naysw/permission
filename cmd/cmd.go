package cmd

import (
	"github.com/naysw/permission/cmd/serve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "permission",
	Short: "permission is a tool for managing permissions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.AddCommand(CreateRoleCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
