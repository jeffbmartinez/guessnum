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
	showGreeting()

	secretNumber := generateSecretNumber()

	win := false

	for guessesLeft := maxGuesses; guessesLeft > 0; guessesLeft-- {
		userGuess := getGuess(guessesLeft)

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

// showGreeting displays a friendly greeting to the user.
func showGreeting() {
	fmt.Println("Welcome to Guess The Number Game!")
	fmt.Println("Try your luck at guessing the correct number!")
}

// generateSecretNumber creates the secret number for the game.
func generateSecretNumber(min int, max int) int {
	return rand.Int()%(max-min+1) + min
}

/*
getGuess retrieves a user's guess from the command line. It allows up to
three attempts to read an integer, then just forces the guess to be zero
if the user can't type an integer.
*/
func getGuess(guessesLeft int) int {
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
