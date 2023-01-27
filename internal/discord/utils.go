package discord

import (
	"fmt"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"github.com/graytonio/ai-dm-prep/internal/generators"
)

func longRunResponse(s *discordgo.Session, i *discordgo.Interaction) error {
	return s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Zolton is Contemplating Your Request...",
		},
	})
}

var emptyString string = ""

func editResponseEmbed(s *discordgo.Session, i *discordgo.Interaction, embed *discordgo.MessageEmbed) error {
	_, err := s.InteractionResponseEdit(i, &discordgo.WebhookEdit{
		Content: &emptyString,
		Embeds:  &[]*discordgo.MessageEmbed{embed},
	})
	return err
}

func errorEmbed(err error) *discordgo.MessageEmbed {
	return embed.NewErrorEmbed("Problem Generating Response", fmt.Sprintf("Looks like Zoltan had a problem generating your response: %v", err))
}

func itemEmbed(item *generators.ItemResponse) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Title:       item.ItemName,
		Description: item.ItemDescription,
		Color:       rand.Intn(16777214) + 1,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Type", Value: cases.Title(language.English).String(item.ItemType), Inline: true},
			{Name: "Rarity", Value: cases.Title(language.English).String(item.ItemRarity), Inline: true},
		},
	}
}
