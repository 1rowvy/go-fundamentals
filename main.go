package main

import (
	"fmt"
)

func main() {
	var x, total int

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		total += i
	}

	fmt.Println(total)
}
