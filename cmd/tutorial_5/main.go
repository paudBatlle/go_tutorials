package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

// Structs and Interfaces
func main() {
	// var myEngine gasEngine // non initialized values = 0
	// fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{25, 15, owner{"Pau"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.name)

}
