package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello", "World", "!")

	fmt.Println("The time rn is ", time.Now())

	fmt.Println("Today is ", time.Now().Weekday())

	x := time.Now()
	fmt.Print(x)
	// fmt.Println("The Day after 3 days is %v", )

}
