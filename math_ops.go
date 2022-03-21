package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

// Go does NOT support function overloading

// func add(x int, y int, z int) int {
// 	return x + y + z
// }

func main() {

	fmt.Println("Value of pi", math.Pi)

	fmt.Println(math.Abs(-25))

	fmt.Println("sqrt of 49 = ", math.Sqrt(49))

	// fmt.Println((add(1, 2, 3)))

	fmt.Println("Adding 2 and 4 = ", add(2, 4))
}
