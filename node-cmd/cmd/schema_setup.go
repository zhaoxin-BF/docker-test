package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhaoxin-BF/docker-test/node-cmd/install/setup"
)

var schemaCreateOptions = struct {
	DatabaseDriver string
	DSN            string
}{}

var schemaCreateCommand = &cobra.Command{
	Use:   "setup",
	Short: "setup schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("setup schema")
		fmt.Println("--------------: ", schemaCreateOptions)

		setup.SetUp()
		return nil
	},
}

func init() {
	flags := schemaCreateCommand.Flags()
	flags.StringVarP(&schemaCreateOptions.DatabaseDriver, "db-driver", "", "mysql", "The driver of database in ['mysql']")
	flags.StringVarP(&schemaCreateOptions.DSN, "dsn", "", "root:123456@tcp(localhost)/resource_nodes", "The dsn of database")

	schemaCommand.AddCommand(schemaCreateCommand)
}
