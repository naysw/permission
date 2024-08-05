package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var CreateRoleCmd = &cobra.Command{
	Use:   "create-role",
	Short: "Create a new role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-role called")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing role name")
		}

		return nil
	},
}
