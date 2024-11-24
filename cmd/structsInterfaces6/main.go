package main

import "fmt"

type owner struct {
	name string
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   // same as ownerInfo owner
	int     // field is called int and is of type int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
	int
}

func (e gasEngine) milesLeft() uint8 { // this is a method!!!!
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 { // this is a method!!!!
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func main() {
	var engine1 gasEngine = gasEngine{mpg: 1, gallons: 2}
	fmt.Println(engine1.mpg, engine1.gallons, engine1.owner, engine1.int)

	var engine2 electricEngine = electricEngine{mpkwh: 25, kwh: 200}

	fmt.Printf("can make it: %v\n", canMakeIt(engine1, 20))
	fmt.Printf("can make it: %v\n", canMakeIt(engine2, 20))
}

func canMakeIt(e engine, miles uint8) bool {
	return miles < e.milesLeft()
}
