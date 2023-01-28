package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/graytonio/ai-dm-prep/internal/generators"
	"github.com/sirupsen/logrus"
)

func init() {
	commands = append(commands, &discordgo.ApplicationCommand{Name: "gen-npc", Description: "Generate a new NPC", Options: []*discordgo.ApplicationCommandOption{
		{
			Type: discordgo.ApplicationCommandOptionString,
			Name: "npc-gender",
			Description: "Gender of NPC",
		},
		{
			Type: discordgo.ApplicationCommandOptionString,
			Name: "npc-race",
			Description: "Race of NPC",
		},
		{
			Type: discordgo.ApplicationCommandOptionString,
			Name: "npc-job",
			Description: "Job of NPC",
		},
	}})
	commandHandlers["gen-npc"] = getNPCHandler
}

func getNPCHandler(s *discordgo.Session, i *discordgo.InteractionCreate, log *logrus.Entry) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	npcGender := optionMap["npc-gender"]
	npcRace := optionMap["npc-race"]
	npcJob := optionMap["npc-job"]

	log.WithFields(logrus.Fields{
		"npc-gender": npcGender,
		"npc-race": npcRace,
		"npc-job": npcJob,
	}).Info("Generating DND NPC")

	err := longRunResponse(s, i.Interaction)
	if err != nil {
		log.Errorf("Error Sending Response: %v", err)
		return
	}

	npc, err := generators.GenerateNPC(npcGender.StringValue(), npcRace.StringValue(), npcJob.StringValue())
	if err != nil {
		log.Errorf("Error Generating Item: %v", err)
		editResponseEmbed(s, i.Interaction, errorEmbed(err))
	}

	log.Info("Seccussfully Generated DND NPC")
	editResponseEmbed(s, i.Interaction, npcEmbed(npc))
}