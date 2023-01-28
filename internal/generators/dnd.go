package generators

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"

	_ "github.com/joho/godotenv/autoload"
)

const ItemGenerator = `Generate a magical %s of %s rarity for use in a Dungeons & Dragons 5th Edition adventure in this format. {"itemName": "","itemDescription": "", "itemRarity": "", "itemType": ""}`
const NPCGenerator = `Generate a %s %s npc %s for use in a Dungeons & Dragons 5th Edition adventure in this format. {"npcName": "","npcDescription": "", "npcClass": "", "npcAlignment": "", "npcStats": { "STR":, "DEX":, "CON":, "INT":, "WIS":, "CHA": }}`

type Item struct {
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
	Type        string `json:"itemType"`
	Rarity      string `json:"itemRarity"`
}

func (item *Item) String() string {
	return fmt.Sprintf("%s: %s", item.Name, item.Description)
}

type NPC struct {
	Name        string `json:"npcName"`
	Description string `json:"npcDescription"`
	Class       string `json:"npcClass"`
	Alignment   string `json:"npcAlignment"`
	Stats       struct {
		Strength     int `json:"STR"`
		Dexterity    int `json:"DEX"`
		Constitution int `json:"CON"`
		Intelligence int `json:"INT"`
		Wisdom       int `json:"WIS"`
		Charisma     int `json:"CHA"`
	} `json:"npcStats"`
}

func (n NPC) String() string {
    return fmt.Sprintf("Name: %s, Description: %s, Class: %s, Alignment: %s, Stats: { Strength: %d, Dexterity: %d, Constitution: %d, Intelligence: %d, Wisdom: %d, Charisma: %d }", n.Name, n.Description, n.Class, n.Alignment, n.Stats.Strength, n.Stats.Dexterity, n.Stats.Constitution, n.Stats.Intelligence, n.Stats.Wisdom, n.Stats.Charisma)
}


var aiClient *gogpt.Client
var ctx context.Context

func init() {
	aiClient = gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx = context.Background()
}

func GenerateNPC(npcGender string, npcRace string, npcJob string) (*NPC, error) {
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Prompt:           fmt.Sprintf(NPCGenerator, npcGender, npcRace, npcJob),
		Temperature:      0,
		BestOf:           1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		TopP:             1,
	}
	resp, err := aiClient.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	var npc NPC
	err = json.Unmarshal([]byte(resp.Choices[0].Text), &npc)
	if err != nil {
		return nil, err
	}

	return &npc, nil
}

func GenerateItem(itemType string, itemRarity string) (*Item, error) {
	if itemType == "" {
		itemType = "item"
	}

	if itemRarity == "" {
		itemRarity = "uncommon"
	}

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Prompt:           fmt.Sprintf(ItemGenerator, itemType, itemRarity),
		Temperature:      0,
		BestOf:           1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		TopP:             1,
	}
	resp, err := aiClient.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	var item Item
	err = json.Unmarshal([]byte(resp.Choices[0].Text), &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
