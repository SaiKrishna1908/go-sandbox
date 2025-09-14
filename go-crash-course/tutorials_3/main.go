package main

import (
	"errors"
	"fmt"
)

func main() {
	// arrays()
	// operationsOnArrays()
	// maps()
	loops()

}

func operationsOnArrays() {
	arr := []int{2, 3, 4, 5, 6, 7, 8}
	arr, err := deleteInArrayByIndex(arr, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v\n", arr)
	}

	arr, isFound := deleteInArrayByValue(arr, 2)
	if isFound {
		fmt.Printf("%v\n", arr)
	} else {
		fmt.Println("Element not found")
	}
}

func arrays() {
	// array declaration
	var arr [3]int32
	arr[1] = 234
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[0:3])

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	// int array fixed size
	var newArr [3]int32 = [3]int32{1, 2, 3}

	fmt.Println(newArr[0])

	newArr1 := [3]int32{1, 2, 3}
	fmt.Println(newArr1)

	newArr2 := [...]int32{1, 2, 3, 4}
	fmt.Println(newArr2)

	// int slice dynamic size
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)

	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7, 8, 9)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	var newIntSlice = []int32{2, 3, 4}
	intSlice = append(intSlice, newIntSlice...)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	// using make to create a int slice. make(type, size, capacity)
	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice3), cap(intSlice3))
}

func maps() {
	// declaration
	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["bod"] = 3
	myMap["Sam"] = 5
	myMap["Gwen"] = 10
	fmt.Printf("%v\n", myMap)

	// declaration in-place

	var inPlaceMap map[string]int = map[string]int{"Adam": 23, "Sarah": 30}
	inPlaceMap["gwent"] = 10
	fmt.Println(inPlaceMap)

	// get the value using key
	age, ok := inPlaceMap["Adam"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Invalid Name")
	}

	// delete an element from map
	delete(inPlaceMap, "Adam")
	fmt.Println(inPlaceMap["Adam"])

	for k, v := range myMap {
		fmt.Printf("%v : %v \n", k, v)
	}
}

func deleteInArrayByIndex(arr []int, index int) ([]int, error) {
	var err error
	if index >= len(arr) {
		err = errors.New("index out of bound")
		return arr, err
	}

	return append(arr[0:index], arr[index+1:]...), nil
}

func deleteInArrayByValue(arr []int, value int) ([]int, bool) {
	res := make([]int, 0)
	isFound := false
	for _, v := range arr {
		if v != value {
			res = append(res, v)
		} else {
			isFound = true
		}
	}

	return res, isFound
}

func loops() {
	i := 10
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	i = 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
}
