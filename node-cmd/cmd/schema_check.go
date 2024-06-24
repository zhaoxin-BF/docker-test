package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var schemaCreateDBOptions = struct {
	DatabaseDriver string
	DSN            string
	dryRun         bool
}{}

var schemaCreateDBCommand = &cobra.Command{
	Use:   "check",
	Short: "check schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("check schema")
		fmt.Println("--------------: ", schemaCreateDBOptions)
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("Current directory:", dir)
		return nil
	},
}

func init() {
	flags := schemaCreateDBCommand.Flags()
	flags.StringVarP(&schemaCreateDBOptions.DatabaseDriver, "db-driver", "", "mysql", "The driver of database in ['mysql']")
	flags.StringVarP(&schemaCreateDBOptions.DSN, "dsn", "", "root:123456@tcp(localhost)/resource_nodes", "The dsn of database")
	flags.BoolVarP(&schemaCreateDBOptions.dryRun, "dry-run", "", false, "Only show sql, nothing be executed")

	schemaCommand.AddCommand(schemaCreateDBCommand)
}
