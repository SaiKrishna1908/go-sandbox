package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()
	printMeWithArgs("Hello World")
	res, rem, err := intDivision(6, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res, rem)
	}
}

func printMe() {
	fmt.Println("Hello world!!")
}

func printMeWithArgs(args string) {
	fmt.Println(args)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var reminder = numerator % denominator
	return result, reminder, nil
}
