package cmd

import (
	"fmt"

	"github.com/graytonio/ai-dm-prep/internal/generators"
	"github.com/spf13/cobra"
)

var genItemType string
var genItemRarity string

func init() {
	genItemCmd.PersistentFlags().StringVar(&genItemType, "type", "", "Type of item to generate (Ex: Sword, Cloak, Ring)")
	genItemCmd.PersistentFlags().StringVar(&genItemRarity, "rarity", "", "Rarity of item to generate (Ex: Uncommon, Rare)")
	genCmd.AddCommand(genItemCmd)
}

var genItemCmd = &cobra.Command{
	Use: "item",
	Short: "Generate a DND Item",
	RunE: func(cmd *cobra.Command, args []string) error {
		item, err := generators.GenerateItem(genItemType, genItemRarity)
		if err != nil {
			return err
		}

		fmt.Println(item)
		return nil
	},
}