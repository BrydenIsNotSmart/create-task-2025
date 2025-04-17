package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func AskToPlay() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Want to play a game of Rock, Paper, and Scissors?: ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("yes", text) == 0 {
			break
		} else if strings.Compare("no", text) == 0 {
			fmt.Println("Okay then, see you some other time.")
			os.Exit(1)
		} else {
			fmt.Println("Please only respond with either \"yes\" or \"no\".")
			continue
		}
	}
}

func AskForName(promt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promt + ": ")

	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	if strings.Compare("", name) == 0 {
		name = "Player"
	}

	return name
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// 0 = CPU | 1 = Human
func AskCPUOrPlayer() int {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Want to play with me (CPU), or another player (Human)?: ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("you", text) == 0 {
			return 0
		} else if strings.Compare("player", text) == 0 {
			return 1
		} else {
			fmt.Println("Please only respond with either \"you\" or \"player\".")
			continue
		}
	}
}

func AskForChoice(playerName string) string {
	ClearTerminal()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s, choose Rock, Paper, or Scissors: ", playerName)
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.TrimSpace(text))

		if text == "rock" || text == "paper" || text == "scissors" {
			return text
		}
		fmt.Println("Invalid choice. Try again.")
	}
}

func GetCPUChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

func DetermineWinner(p1Choice, p2Choice string) int {
	if p1Choice == p2Choice {
		return 0
	}
	switch p1Choice {
	case "rock":
		if p2Choice == "scissors" {
			return 1
		}
	case "paper":
		if p2Choice == "rock" {
			return 1
		}
	case "scissors":
		if p2Choice == "paper" {
			return 1
		}
	}
	return 2
}
