package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	//Commands
	helloCommand := flag.NewFlagSet(HELLO, flag.ExitOnError)
	fizzBuzzCommand := flag.NewFlagSet(FIZZBUZZ, flag.ExitOnError)

	//Options
	helloNameString := helloCommand.String("name", "", "Your name. (Required)")
	fizzBuzzNumbersString := fizzBuzzCommand.String("numbers", "", "Comma separated string of number. (Required)")

	if len(os.Args) < 2 {
		flag.PrintDefaults()

		fmt.Println("Commands:")
		fmt.Println("   "+HELLO+"\t output your name.")
		fmt.Println("   "+FIZZBUZZ+"\t for multiples of three print 'Fizz' and for multiples of five print 'Buzz'.")
		return
	}

	switch os.Args[1] {
	case HELLO:
		helloCommand.Parse(os.Args[2:])
	case FIZZBUZZ:
		fizzBuzzCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return
	}

	if helloCommand.Parsed() {
		if *helloNameString == "" {
			helloCommand.PrintDefaults()
		}else {
			hello(*helloNameString)
		}
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