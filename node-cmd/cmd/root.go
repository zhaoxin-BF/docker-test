package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhaoxin-BF/docker-test/node-cmd/utils"
)

type Options struct {
	EverAIHomePath string // for set reverse-proxy socket path
}

var options = Options{}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("resource-node")
		fmt.Printf("options: %+v\n", options)
		utils.SetEveraiNodeHome(options.EverAIHomePath)
		fmt.Println(utils.DBPath)
		fmt.Println(utils.VolumeRoot)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&options.EverAIHomePath, "everai-home-path", "", "/data/everai/node", "everai home path")
}
