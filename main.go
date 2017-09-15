package main

import (
	"flag"
	"os"
	"fmt"
)

const COMMAND_HELLO  = "hello";
const COMMAND_TIME = "time";
const COMMAND_GUESS_NUMBER = "guessnumber";
const COMMAND_FIZZ_BUZZ = "fizzbuzz";

func main() {
	//Commands
	helloCommand := flag.NewFlagSet(COMMAND_HELLO, flag.ExitOnError)
	timeCommand := flag.NewFlagSet(COMMAND_TIME, flag.ExitOnError)
	guessNumberCommand := flag.NewFlagSet(COMMAND_GUESS_NUMBER, flag.ExitOnError)
	fizzBuzzCommand := flag.NewFlagSet(COMMAND_FIZZ_BUZZ, flag.ExitOnError)

	//Options
	helloNameString := helloCommand.String("name", "", "Your name. (Required)")
	fizzBuzzNumbersString := fizzBuzzCommand.String("numbers", "", "Comma separated string of number. (Required)")

	if len(os.Args) < 2 {
		flag.PrintDefaults()

		fmt.Println("Commands:")
		fmt.Println("   "+COMMAND_HELLO+"\t say hello.")
		fmt.Println("   "+COMMAND_TIME+"\t\t output the current time.")
		fmt.Println("   "+COMMAND_GUESS_NUMBER+"\t try to guess the rand number.")
		fmt.Println("   "+COMMAND_FIZZ_BUZZ+"\t for multiples of three print 'Fizz' and for multiples of five print 'Buzz'.")
		return
	}

	switch os.Args[1] {
	case COMMAND_HELLO:
		helloCommand.Parse(os.Args[2:])
	case COMMAND_TIME:
		timeCommand.Parse(os.Args[2:])
	case COMMAND_GUESS_NUMBER:
		guessNumberCommand.Parse(os.Args[2:])
	case COMMAND_FIZZ_BUZZ:
		fizzBuzzCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return
	}

	//Hello
	if helloCommand.Parsed() {
		if *helloNameString == "" {
			helloCommand.PrintDefaults()
		}else {
			hello := Hello{Name:*helloNameString}
			hello.hello();
		}
	}

	//Time
	if timeCommand.Parsed() {
		timestamp()
	}

	//Guess Number
	if guessNumberCommand.Parsed() {
		guessNumber := GuessNumber{Guesses: 0, Number: 0}
		guessNumber.setNumber()
		guessNumber.playGuessNumber();
	}

	//FizzBuzz
	if fizzBuzzCommand.Parsed() {
		if *fizzBuzzNumbersString == "" {
			fizzBuzzCommand.PrintDefaults()
		}else {
			playFizzBuzz(*fizzBuzzNumbersString)
		}
	}
}