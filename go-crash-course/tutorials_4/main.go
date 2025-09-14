package main

import (
	"fmt"
	"strings"
)

func main() {
	stringsInGo()
}

func runes() {
	var myString = "resume"
	var indexed = myString[0]
	fmt.Printf("%v %v\n", myString, indexed)
	// print type of indexed
	fmt.Printf("%T\n", indexed)

	// iterate through index
	for i, v := range myString {
		fmt.Printf("%v, %v\n", i, v)
	}

	var runeArr = []rune("resume")
	var runeIndexed = runeArr[0]
	fmt.Printf("%v", runeIndexed)
}

func stringsInGo() {
	var strSlice = []string{"h", "e", "l", "l", "0"}

	// immutable
	var str = ""
	for i := range strSlice {
		str += strSlice[i]
	}

	var stringBuilder strings.Builder
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}

	fmt.Printf("\n%v", stringBuilder.String())
}
