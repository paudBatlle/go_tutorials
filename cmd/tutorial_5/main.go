package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electrincEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electrincEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it !")
	} else {
		fmt.Println("Need to refuel first")
	}
}

// Structs and Interfaces
func main() {
	// var myEngine gasEngine // non initialized values = 0
	// fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{25, 15, owner{"Pau"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.name)
	//fmt.Printf("Total miles left in tank: %v", myEngine2.milesLeft())
	canMakeIt(myEngine2, 50)
}
