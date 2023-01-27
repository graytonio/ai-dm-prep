package main

import (
	"fmt"
	"os"

	"github.com/graytonio/ai-dm-prep/internal/discord"
	"github.com/graytonio/ai-dm-prep/internal/generators"
)

func cliUsage() {
	inputType := os.Args[1]
	inputRarity := os.Args[2]

	item, err := generators.GenerateItem(inputType, inputRarity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(item)
}

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		logrus.Fatalln("Error loading .env file", err)
// 	}
// }

func main() {
	discord.StartServer()
}
