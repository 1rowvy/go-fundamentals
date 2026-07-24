package main

import "fmt"

func main() {
	var max int

	for {
		var a int
		if _, err := fmt.Scan(&a); err != nil {
			break
		}

		if a > max {
			max = a
		}
	}

	fmt.Print(max)

}
