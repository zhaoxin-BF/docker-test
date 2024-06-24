package cmd

import (
	"github.com/spf13/cobra"
)

var schemaCommand = &cobra.Command{
	Use:   "schema",
	Short: "Schema management",
}

func init() {
	rootCmd.AddCommand(schemaCommand)
}
