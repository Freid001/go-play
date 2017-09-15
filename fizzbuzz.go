package main

import (
	"strings"
	"fmt"
	"strconv"
	"os"
)

type FizzBuzz struct{
	Number int64
	Fizz bool
	Buzz bool
}

func playFizzBuzz(numbersString string){
	//Split numbers string into array
	numbers := strings.Split(numbersString, ",")

	//Iterate over array
	for i := 0; i < len(numbers); i += 1 {

		//Do we have a value?
		if(numbers[i] == ""){
			continue
		}

		//Convert number to int
		number, err := strconv.ParseInt(numbers[i], 10, 0)

		//Output error
		if(err != nil) {
			fmt.Fprintf(os.Stderr, "Encountered error: %v", err)
			fmt.Println();
			break;
		}

		//Is this a fizz buzz number?
		fizzBuzz := isNumberAFizzBuzz(number)
		defer fmt.Println();

		//Output FIZZ
		if (fizzBuzz.Fizz) {
			defer fmt.Print("[FIZZ]");
		}

		//Output BUZZ
		if (fizzBuzz.Buzz) {
			defer fmt.Print("[BUZZ]");
		}

		//Output our number
		defer fmt.Print("# " + strconv.FormatInt(fizzBuzz.Number, 10) + " ");
	}
}

func isNumberAFizzBuzz(number int64) FizzBuzz {
	if(number % 3 == 0 && number % 5 == 0){
		return FizzBuzz{Number:number,Fizz:true,Buzz:true}
	}else if(number % 3 == 0){
		return FizzBuzz{Number:number,Fizz:true,Buzz:false}
	}else if(number % 5 == 0){
		return FizzBuzz{Number:number,Fizz:false,Buzz:true}
	}

	return FizzBuzz{Number:number,Fizz:false,Buzz:false}
}