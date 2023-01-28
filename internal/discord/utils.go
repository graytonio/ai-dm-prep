package discord

import (
	"fmt"
	"math/rand"

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

func randomEmbedColor() int {
	return rand.Intn(16777214) + 1
}

func itemEmbed(item *generators.Item) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
        Title:       item.Name,
        Description: item.Description,
		Color: randomEmbedColor(),
        Fields: []*discordgo.MessageEmbedField{
            &discordgo.MessageEmbedField{
                Name:  "Type",
                Value: item.Type,
				Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name:  "Rarity",
                Value: item.Rarity,
				Inline: true,
            },
        },
    }
}

func npcEmbed(npc *generators.NPC) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
        Title:       npc.Name,
        Description: npc.Description,
        Fields: []*discordgo.MessageEmbedField{
            &discordgo.MessageEmbedField{
                Name:  "Class",
                Value: npc.Class,
            },
            &discordgo.MessageEmbedField{
                Name:  "Alignment",
                Value: npc.Alignment,
            },
            &discordgo.MessageEmbedField{
                Name:  "STR",
                Value: fmt.Sprintf("%d", npc.Stats.Strength),
				Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name:  "DEX",
                Value: fmt.Sprintf("%d", npc.Stats.Dexterity),
				Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name:  "CON",
                Value: fmt.Sprintf("%d", npc.Stats.Constitution),
				Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name:  "INT",
                Value: fmt.Sprintf("%d", npc.Stats.Intelligence),
				Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name: "WIS",
				Value: fmt.Sprintf("%d", npc.Stats.Wisdom),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
                Name: "CHA",
				Value: fmt.Sprintf("%d", npc.Stats.Charisma),
				Inline: true,
			},
		},
	}
}