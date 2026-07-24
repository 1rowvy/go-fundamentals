package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	switch {
	case x%15 == 0:
		fmt.Print("FizzBuzz")
	case x%3 == 0:
		fmt.Print("Fizz")
	case x%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Println(x)
	}
}
