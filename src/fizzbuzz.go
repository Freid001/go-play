package main

import (
	"strings"
	"strconv"
	"fmt"
)

const FIZZBUZZ = "fizzbuzz"

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

		//Convert number to int
		number, err := strconv.ParseInt(numbers[i], 10, 0)

		fmt.Print(err)

		//Is this a fizz buzz number?
		fizzBuzz := isNumberAFizzBuzz(number)

		//fmt.Print(fizzBuzz)

		//Output our number
		fmt.Print("# " + strconv.FormatInt(fizzBuzz.Number, 10) + " ");

		//Output FIZZ
		if(fizzBuzz.Fizz){
			fmt.Print("[FIZZ]");
		}

		//Output BUZZ
		if(fizzBuzz.Buzz){
			fmt.Print("[BUZZ]");
		}

		fmt.Println();
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