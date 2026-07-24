package main

import "fmt"

func main() {

	var x int

	fmt.Scan(&x)

	fmt.Println(square(x))

}

func square(n int) int {
	return n * n
}
