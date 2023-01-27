package discord

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

var session *discordgo.Session
var commands = []*discordgo.ApplicationCommand{}
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, log *log.Entry){}

var AppID string = "1068650403715100753"
var GuildID string = "1068650327462662219"

func init() {
	var err error
	session, err = discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("DISCORD_BOT_TOKEN")))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot Ready")
	})

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			handleApplicationCommand(s, i)
		}
	})
}

func handleApplicationCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		logger := log.WithField("interaction-id", i.ID).WithField("command", i.ApplicationCommandData().Name).WithField("guild", i.GuildID)
		logger.Info("Executing Command")
		h(s, i, logger)
	}
}

func StartServer() {
	_, err := session.ApplicationCommandBulkOverwrite(os.Getenv("DISCORD_APP_ID"), os.Getenv("DISCORD_TEST_SERVER_ID"), commands)
	if err != nil {
		log.Fatalf("Cannot create application commands: %v", err)
	}

	err = session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful Shutdown")
}
