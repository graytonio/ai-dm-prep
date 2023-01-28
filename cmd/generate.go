package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use: "gen",
	Short: "Use a generator in the cli",
}