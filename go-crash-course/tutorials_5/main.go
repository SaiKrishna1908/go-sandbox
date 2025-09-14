package main

import "fmt"

type Position struct {
	x int
	y int
}

func (p *Position) addOne() {
	p.x += 1
	p.y += 1
}

func (p *Position) subOne() {
	p.x -= 1
	p.y -= 1
}

func structsInGo() {
	var pos Position = Position{y: 5, x: 10}
	fmt.Printf("%v\n", pos)

	var personStruct = struct {
		name string
		age  uint8
	}{"bob", 23}
	fmt.Printf("%v\n", personStruct)

	var p Position = Position{0, 0}
	fmt.Printf("%v\n", p)
	p.addOne()
	fmt.Printf("%v\n", p)
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it!")
	} else {
		fmt.Println("Need fuel")
	}
}

func main() {
	canMakeIt(gasEngine{9, 9}, 10)
	canMakeIt(electricEngine{10, 2}, 22)
}
