package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var intNum int
	fmt.Println(intNum)
	fmt.Println("Hello World!!")

	var floatNum float32
	fmt.Println(floatNum)

	// casting uint8 to int
	var rgb uint8
	fmt.Println(rgb)
	fmt.Println(int(rgb))

	var myString string = "Hello World!!"
	fmt.Println(myString)

	// get string length
	fmt.Println(utf8.RuneCountInString(myString))

	// rune's are characters
	var myRune rune = 'a'
	// outputs 97
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	// type inference in golang

	intNum5 := 5
	fmt.Println(intNum5)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	var3, var4 := '1', "2"

	fmt.Println(var3, var4)

	const myConst string = "this is fixed!!"
	fmt.Println(myConst)

	const pi = 3.142
	fmt.Println(pi)
}
