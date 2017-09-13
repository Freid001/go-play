package main

import (
	"flag"
	"fmt"
	_ "github.com/urfave/cli"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)




	switch arg{
	case FIZZBUZZ:
		//fmt.Print("Enter numbers: ")
		////numbers, _ := reader.ReadString('');

		//Run fizz buzz
		playFizzBuzz(numbers)
	default:
		fmt.Println("Commands:");
		fmt.Println("");
		fmt.Println("    fizzbuzz \t for multiples of three print 'Fizz' and for multiples of five print 'Buzz'.");
	}
}