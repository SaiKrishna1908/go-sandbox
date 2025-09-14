package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type absBrakingSystem struct {
	name         string
	manufacturer string
}

type nonAbsBrakingSystem struct {
	name         string
	manufacturer string
}

type Car[T gasEngine | electricEngine, F absBrakingSystem | nonAbsBrakingSystem] struct {
	carMake       string
	carModel      string
	engine        T
	brakingSystem F
}

func createCar[T gasEngine | electricEngine, F absBrakingSystem | nonAbsBrakingSystem](engine T, braking F) Car[T, F] {
	return Car[T, F]{
		carMake:       "",
		carModel:      "",
		engine:        engine,
		brakingSystem: braking,
	}
}

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(intSlice)
	// fmt.Println(sumIntSlice(intSlice))
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(float32Slice)
	// fmt.Println(sumFloat32Slice(float32Slice))
	fmt.Println(sumSlice[float32](float32Slice))

	var float64Slice = []float64{1, 2, 3}
	fmt.Println(float64Slice)
	// fmt.Println(sumFloat64Slice(float64Slice))
	fmt.Println(sumSlice[float64](float64Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
