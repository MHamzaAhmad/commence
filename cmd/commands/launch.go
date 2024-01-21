package commands

import (
	"github.com/MhamzaAhmad/commence/executer"
	"github.com/spf13/cobra"
)

func CreateLaunchCommand() *cobra.Command {
	return &cobra.Command{
	Use:   "launch",
	Short: "Launches the saved named sequence from storage",
	Long: `Launches the saved named sequence from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		executer.Execute(name)
	},
}
}
