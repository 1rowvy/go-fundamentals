package main

import (
	"fmt"
)

func main() {
	var (
		name string
		age  int
	)

	fmt.Scan(&name, &age)

	fmt.Printf("Hi, %s! You are %d years old.", name, age)
}
