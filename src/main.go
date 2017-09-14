package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	//Commands
	fizzBuzzCommand := flag.NewFlagSet(FIZZBUZZ, flag.ExitOnError)

	//Options
	fizzBuzzNumbersString := fizzBuzzCommand.String("numbers", "", "Comma separated string of number. (Required)")

	if len(os.Args) < 2 {
		flag.PrintDefaults()

		fmt.Println("Commands:")
		fmt.Println("   "+FIZZBUZZ+"\t for multiples of three print 'Fizz' and for multiples of five print 'Buzz'.")
		return
	}

	switch os.Args[1] {
	case FIZZBUZZ:
		fizzBuzzCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return
	}

	if fizzBuzzCommand.Parsed() {
		if *fizzBuzzNumbersString == "" {
			fizzBuzzCommand.PrintDefaults()
		}else {
			playFizzBuzz(*fizzBuzzNumbersString)
		}
	}
}
