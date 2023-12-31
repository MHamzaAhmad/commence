package cmd

import (
	"os"

	"github.com/MhamzaAhmad/commence/cmd/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commence",
	Short: "A tool to help you launch your project on everyday bases",
	Long: `A cli tool to help launch multiple utilities that are needed to start developmnent of you project. I requires YAML file write in easy language containing the steps of the launch sequence`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.CreateLoadCommand())
}


