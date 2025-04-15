package main

import (
	"fmt"
	"time"

	"github.com/BrydenIsNotSmart/create-task-2025/utils"
)

func main() {
	utils.AskToPlay()
	playerName := utils.AskForName("What should I call you?")
	humanOrCPU := utils.AskCPUOrPlayer()
	fmt.Println(humanOrCPU)
	fmt.Printf("Alright %v, let's get started!\n", playerName)

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	utils.ClearTerminal()

}
