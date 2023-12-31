package commands

import (
	"github.com/MhamzaAhmad/commence/parser"
	"github.com/spf13/cobra"
)

func CreateLoadCommand() *cobra.Command {
	return &cobra.Command{
	Use:   "load",
	Short: "Loads the sequence of commands from the YAML file",
	Long: `Loads the sequence of commands from the YAML file`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		parser.Parse(path)
	},
}
}
