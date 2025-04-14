package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskToPlay() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Want to play a game of Rock, Paper, and Scissors with me?: ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("yes", text) == 0 {
			fmt.Println("Let's get started.")
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
