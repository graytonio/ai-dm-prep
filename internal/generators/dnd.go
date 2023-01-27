package generators

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"

	_ "github.com/joho/godotenv/autoload"
)
const ItemGenerator = `Generate a magical %s of %s rarity for use in a Dungeons & Dragons 5th Edition adventure in this format. {"itemName": "","itemDescription": ""}`

type ItemResponse struct {
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	ItemType        string
	ItemRarity      string
}

func (item *ItemResponse) String() string {
	return fmt.Sprintf("%s: %s", item.ItemName, item.ItemDescription)
}

var aiClient *gogpt.Client
var ctx context.Context

func init() {
	aiClient = gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx = context.Background()
}

func GenerateItem(itemType string, rarity string) (*ItemResponse, error) {
	if itemType == "" {
		itemType = "item"
	}

	if rarity == "" {
		rarity = "uncommon"
	}

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Prompt:           fmt.Sprintf(ItemGenerator, itemType, rarity),
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

	var item ItemResponse
	err = json.Unmarshal([]byte(resp.Choices[0].Text), &item)
	if err != nil {
		return nil, err
	}

	item.ItemRarity = rarity
	item.ItemType = itemType

	return &item, nil
}
