package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const minNumber = 1
const maxNumber = 100

const maxGuesses = 8

func main() {
	ShowGreeting()

	secretNumber := GenerateSecretNumber(minNumber, maxNumber)

	win := false

	for guessesLeft := maxGuesses; guessesLeft > 0; guessesLeft-- {
		userGuess := GetGuess(guessesLeft)

		switch {
		case userGuess == secretNumber:
			fmt.Println("Wizard!! That's some next-level mind reading skillz!")
			win = true
		case userGuess < secretNumber:
			fmt.Println("Whoa! Too low!")
		case userGuess > secretNumber:
			fmt.Println("Nice try! Too high!")
		default:
			panic("That's just bananas!")
		}

		if win {
			break
		}
	}

	if !win {
		fmt.Printf("\nThe number was %v, but nice try!\n", secretNumber)
	}
}

// ShowGreeting displays a friendly greeting to the user.
func ShowGreeting() {
	fmt.Println("Welcome to Guess The Number Game!")
	fmt.Println("Try your luck at guessing the correct number!")
}

// GenerateSecretNumber creates the secret number for the game.
func GenerateSecretNumber(min int, max int) int {
	return rand.Int()%(max-min+1) + min
}

/*
GetGuess retrieves a user's guess from the command line. If it's unable to
read an integer from what the user types, it just pretends the user entered
zero. But don't do that, it's a toy program that assumes good input. Weird
things happen if you enter non integer input.
*/
func GetGuess(guessesLeft int) int {
	guessWord := "guesses"
	if guessesLeft == 1 {
		guessWord = "guess"
	}

	fmt.Printf("===== You have %v %v left! =====\n", guessesLeft, guessWord)
	fmt.Printf("The integer is between %v and %v (inclusive)\n", minNumber, maxNumber)
	fmt.Printf("Enter your guess: ")

	guess := int(0)
	_, err := fmt.Scanln(&guess)
	if err != nil {
		fmt.Println("Do you even math bro? That's not even an integer!")
		fmt.Println("Obviously you meant zero so we'll go with that")
		guess = 0
	}

	return guess
}
