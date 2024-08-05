package serve

import (
	"fmt"

	"github.com/naysw/permission/api/rest"
	"github.com/naysw/permission/internal/core"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the permission server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	fmt.Println("Starting server on port 9000")
	if err := rest.StartServer(core.NewApp(), rest.WithPort("9000")); err != nil {
		panic(err)
	}
}
