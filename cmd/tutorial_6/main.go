package main

import "fmt"

// pointers
func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	*p = 10 // assign a value to the memory that p points to
	p = &i  // assign p the memory adress of i
	*p = 1  // now we change the value of i
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v \n", i)

	var slice = []int32{1, 2, 3}
	var SliceCopy = slice
	SliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(SliceCopy) // both variables refer to the same data (copying pointers)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\n The result is: %v", result)
	fmt.Printf("\n The value of thing1 is: %v", thing1)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
