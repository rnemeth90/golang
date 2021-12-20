package main

import (
	"fmt"
	"myapp/game"
)

func main() {
	playAgain := true

	for playAgain {
		// game.Play()
		playAgain = game.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}
