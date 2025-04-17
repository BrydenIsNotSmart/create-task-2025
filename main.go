package main

import (
	"fmt"
	"time"

	"github.com/BrydenIsNotSmart/create-task-2025/utils"
)

func main() {
	utils.AskToPlay()
	player1Name := utils.AskForName("What should I call you?")
	humanOrCPU := utils.AskCPUOrPlayer()

	var player2Name string
	if humanOrCPU == 1 {
		player2Name = utils.AskForName("What should I call the second player?")
	} else {
		player2Name = "CPU"
	}

	fmt.Printf("Alright %v and %v, let's get started!\n", player1Name, player2Name)

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	utils.ClearTerminal()

	player1 := utils.Player{Name: player1Name}
	player2 := utils.Player{Name: player2Name}
	game := utils.Game{Player1: player1, Player2: player2, Rounds: 3}

	for round := 1; round <= game.Rounds; round++ {
		fmt.Printf("Round %d!\n", round)
		game.Player1.Choice = utils.AskForChoice(game.Player1.Name)

		if humanOrCPU == 1 {
			game.Player2.Choice = utils.AskForChoice(game.Player2.Name)
		} else {
			game.Player2.Choice = utils.GetCPUChoice()
			fmt.Printf("%s chose %s!\n", game.Player2.Name, game.Player2.Choice)
			time.Sleep(1 * time.Second)
		}

		result := utils.DetermineWinner(game.Player1.Choice, game.Player2.Choice)
		switch result {
		case 0:
			fmt.Println("It's a tie!")
		case 1:
			fmt.Printf("%s wins the round!\n", game.Player1.Name)
			game.Player1.Score++
		case 2:
			fmt.Printf("%s wins the round!\n", game.Player2.Name)
			game.Player2.Score++
		}

		time.Sleep(2 * time.Second)
		utils.ClearTerminal()
	}

	fmt.Println("Final Scores:")
	fmt.Printf("%s: %d\n", game.Player1.Name, game.Player1.Score)
	fmt.Printf("%s: %d\n", game.Player2.Name, game.Player2.Score)

	if game.Player1.Score > game.Player2.Score {
		fmt.Printf("%s wins the game!\n", game.Player1.Name)
	} else if game.Player1.Score < game.Player2.Score {
		fmt.Printf("%s wins the game!\n", game.Player2.Name)
	} else {
		fmt.Println("The game is a tie!")
	}
}
