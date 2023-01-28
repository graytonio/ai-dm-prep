package cmd

import (
	"fmt"

	"github.com/graytonio/ai-dm-prep/internal/generators"
	"github.com/spf13/cobra"
)

var genNPCGender string
var genNPCRace string
var genNPCJob string

func init() {
	genNPCCmd.PersistentFlags().StringVar(&genNPCGender, "gender", "", "Gender of NPC")
	genNPCCmd.PersistentFlags().StringVar(&genNPCRace, "race", "", "Race of NPC")
	genNPCCmd.PersistentFlags().StringVar(&genNPCJob, "job", "", "Job of NPC")
	genCmd.AddCommand(genNPCCmd)
}

var genNPCCmd = &cobra.Command{
	Use: "npc",
	Short: "Generate a DND NPC",
	RunE: func(cmd *cobra.Command, args []string) error {
		npc, err := generators.GenerateNPC(genNPCGender, genNPCRace, genNPCJob)
		if err != nil {
			return err
		}

		fmt.Println(npc)
		return nil
	},
}