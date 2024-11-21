package main

import (
	"fmt"
	"time"
)

// Arrays, Slices, Maps
// Looping control structures

func main() {
	// Arrays

	var intArr [3]int32 // fixedd length and same type
	intArr[1] = 123     // indexable
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Contiguous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)
	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3)
	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr4)

	// Slices
	var intSlice []int32 = []int32{4, 5, 6} // [4,5,6]
	fmt.Println(intSlice)
	fmt.Printf("The length is %v and the capacity is %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // [4,5,6,7,*,*]
	fmt.Println(intSlice)
	fmt.Printf("The length is %v and the capacity is %v", len(intSlice), cap(intSlice))
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // [4,5,6,7,8,9]
	fmt.Println(intSlice)
	var intSlice3 []int32 = make([]int32, 3, 8) //(type, length, capacity)
	fmt.Println(intSlice3)

	// Maps (deictionary {key:value})
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 map[string]int32 = map[string]int32{"one": 1, "two": 2}
	fmt.Println(myMap2)
	fmt.Println(myMap2["one"])
	fmt.Println(myMap2["three"])
	var val, ok = myMap2["one"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Key not found")
	}
	// delete(myMap2, "one")
	// var val2, ok2 = myMap2["one"]
	// if ok2 {
	// 	fmt.Println(val2)
	// } else {
	// 	fmt.Println("Key not found")
	// }

	// Looping control structures
	for num, num2 := range myMap2 {
		fmt.Printf("Number: %v, number 2: %v \n", num, num2) // no order is preserved
	}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int = 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	// Speed test for capacity of slice
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return time.Since(start)
}
