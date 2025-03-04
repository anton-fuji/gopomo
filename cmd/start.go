package cmd

import (
	"gopomo/ui"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a 25-minute Pomodoro session",
	Run: func(cmd *cobra.Command, args []string) {
		ui.RunTimer(25)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
