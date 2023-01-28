package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/graytonio/ai-dm-prep/internal/generators"
	"github.com/sirupsen/logrus"
)

func init() {
	commands = append(commands, &discordgo.ApplicationCommand{Name: "gen-item", Description: "Generate a new item", Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "item-type",
			Description: "Item Type (Example: Sword, Ring, Cloak)",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "item-rarity",
			Description: "Item Rarity (Example: Uncommon, Rare)",
			Required:    true,
		},
	}})
	commandHandlers["gen-item"] = getItemHandler
}

func getItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate, log *logrus.Entry) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	itemType := optionMap["item-type"]
	itemRarity := optionMap["item-rarity"]

	log.WithFields(logrus.Fields{
		"item-type": itemType.StringValue(),
		"item-rarity": itemRarity.StringValue(),
	}).Info("Generating DND Item")

	err := longRunResponse(s, i.Interaction)
	if err != nil {
		log.Errorf("Error Sending Response: %v", err)
		return
	}

	item, err := generators.GenerateItem(itemType.StringValue(), itemRarity.StringValue())
	if err != nil {
		log.Errorf("Error Generating Item: %v", err)
		editResponseEmbed(s, i.Interaction, errorEmbed(err))
	}

	log.Info("Successfully Generated DND Item")
	editResponseEmbed(s, i.Interaction, itemEmbed(item))
}
