package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var IntNum int = 10
	fmt.Println(IntNum)
	var int64num int64 = 100
	fmt.Println(int64num)
	var int16num int16 = 32767
	fmt.Println(int16num)
	var uintNum uint = 100
	fmt.Println(uintNum)
	var floatNum float32 = 3.1415926535
	fmt.Println(floatNum)
	var floatNum64 float64 = 3.1415926535
	fmt.Println(floatNum64)

	// aritmetic operation have to be of the same type
	//Error: var sum int = IntNum + int64num
	var sum int = IntNum + int(int64num)
	fmt.Println(sum)

	var MyString string = "Hello World"
	fmt.Println(MyString)
	var stringg string = `Hello
	world`
	fmt.Println(stringg)

	fmt.Println(len(MyString))
	fmt.Println(utf8.RuneCountInString(MyString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	myvar := "text"
	fmt.Println(myvar)
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst int = 10
	fmt.Println(myConst)
}
