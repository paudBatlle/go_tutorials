package main

import (
	"fmt"
	"strings"
)

// utf8: variable length encoding

func main() {
	var myString = "résumé"
	fmt.Println(myString)
	var indexed = myString[0]
	fmt.Println(indexed)                       // printing an indexed string outputs a number
	fmt.Printf("%v, %T \n ", indexed, indexed) // instead of string --> uint8

	for i, v := range myString {
		fmt.Println(i, v) // é is represented as 2bytes in memory a
	}
	// Most of operations with strings deals with the bytes in meory
	// the best way to solve this problem is by casting to an array of runes aka int32
	var myString2 = []rune("résumé")
	fmt.Println(myString2)
	var indexed2 = myString2[0]
	fmt.Println(indexed2)
	fmt.Printf("%v, %T \n", indexed2, indexed2)

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	// concatenating strings
	var strSlice = []string{"p", "a", "u", "d", "e", "b", "a", "t", "l", "l", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	var strBuilder strings.Builder // create an array in memory
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) //assign the characters to the array
	}
	var catStr2 = strBuilder.String() // create a string from the array
	fmt.Printf("\n%v", catStr2)

}
