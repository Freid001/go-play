package main

import (
	"math/rand"
	"fmt"
	"time"
	"github.com/c-bata/go-prompt"
	"strconv"
	"os"
)

const GuessingChances = 6
const MaxNumber = 19

type GuessNumber struct {
	Guesses int
	Number  int
}

func (guessNumber *GuessNumber) setNumber() {
	if (guessNumber.Number == 0) {
		rand.Seed(time.Now().UnixNano())
		guessNumber.Number = rand.Intn(MaxNumber)
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (guessNumber *GuessNumber) playGuessNumber(){
	for i := 0; i <= GuessingChances-1; i++ {

		//Request a input guess
		fmt.Println("Make a guess")
		guess := prompt.Input("> ", completer)

		//Convert to int
		number, err := strconv.Atoi(guess);

		if(err != nil){
			fmt.Fprintf(os.Stderr, "Encountered error: %v", err)
			fmt.Println();
			break;
		}

		//Is this guess the correct answer?
		if(guessNumber.isAnswer(number)){
			break;
		}
	}
}

func (guessNumber *GuessNumber) isAnswer(guess int) (bool) {
	guessNumber.Guesses++;

	switch {
	case guessNumber.Guesses >= GuessingChances:
		fmt.Println("Better luck next time! Your out of guess!")
		return false;
	case guess <= 0:
		fmt.Printf("Your answer is far too small. The number lies between 1 and %d [Guess: %d of %d]\n", MaxNumber, guessNumber.Guesses, GuessingChances)
		return false;
	case guess > MaxNumber:
		fmt.Printf("Your answer is far too big. The number lies between 1 and %d [Guess: %d of %d]\n", MaxNumber, guessNumber.Guesses, GuessingChances)
		return false;
	case guess < guessNumber.Number:
		fmt.Printf("%d is too low. [Guess: %d of %d]\n", guess, guessNumber.Guesses, GuessingChances)
		return false;
	case guess > guessNumber.Number:
		fmt.Printf("%d is too high. [Guess: %d of %d]\n", guess, guessNumber.Guesses, GuessingChances)
		return false;
	}

	fmt.Printf("Good job! You guessed my number in %d guesses! \n", guessNumber.Guesses)

	return true;
}