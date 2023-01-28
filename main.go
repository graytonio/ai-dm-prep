package main

import (
	"github.com/graytonio/ai-dm-prep/cmd"
)

// func cliUsage() {
// 	inputType := os.Args[1]
// 	inputRarity := os.Args[2]

// 	item, err := generators.GenerateItem(inputType, inputRarity)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(item)
// }

func main() {
	// discord.StartServer()
	cmd.Execute()
}
