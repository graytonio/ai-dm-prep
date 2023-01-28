package cmd

import (
	"github.com/graytonio/ai-dm-prep/internal/discord"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "zoltan-bot",
	Short: "A discord bot for generating AI DND resources",
	Run: func(cmd *cobra.Command, args []string) {
		discord.StartServer()
	},
}

func Execute() error {
	return rootCmd.Execute()
}