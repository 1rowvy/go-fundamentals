package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	increment(&n)

	fmt.Print(n)
}

func increment(n *int) {
	*n++
}
