package main

import "fmt"

func main() {
	pointers()
	sliceCopyByAddress()

	var imutArray = [5]float32{1, 2, 3, 4, 5}
	squareAnArray(&imutArray)
	fmt.Println(imutArray)
}

func squareAnArray(arr *[5]float32) {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
}

func pointers() {
	// initalize a pointer using *
	var p *int32 = new(int32)
	var i int32 = 12

	// dereference to the value of the pointer using *
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

	*p = 10
	fmt.Printf("The value p points to is %v\n", *p)

	// create a pointer from the address of another variable
	p = &i
	fmt.Printf("The value p points to is %v\n", *p)
}

func sliceCopyByAddress() {
	var slice = []int32{1, 2, 3, 4}
	var copySlice = slice
	copySlice[1] = 3

	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", copySlice)
}
