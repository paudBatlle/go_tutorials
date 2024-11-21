package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
	}
	// if 1 ==1 && 2==2 || 3==3{}
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)

	}
	switch remainder {
	case 0:
		fmt.Printf("The result of the integer division was exact")
	case 1, 2:
		fmt.Printf("The remainder was 1 or 2")
	default:
		fmt.Print("The remainder was greater than 2")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
