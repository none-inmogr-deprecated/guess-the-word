package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"guess-the-word/game"
	"guess-the-word/generator"
	"guess-the-word/random"
)

func main() {
	// display message
	fmt.Println("Choose type of words to guess from:")
	fmt.Println("1- Colors")
	fmt.Println("2- Directions")
	fmt.Println("3- Names")
	fmt.Print("Enter text: ")

	// take input
	reader := bufio.NewReader(os.Stdin)
	guessGame, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	guessGame = strings.TrimSuffix(guessGame, "\n")
	guessGame = strings.TrimSuffix(guessGame, string(13))

	words := generator.Generate(guessGame)
	shuffled := random.Shuffle(words)

	score := game.Start(words, shuffled)

	fmt.Print("\n\n")
	fmt.Println("Great you guessed it all!")
	fmt.Printf("Your score is %.2f\n", score)
}