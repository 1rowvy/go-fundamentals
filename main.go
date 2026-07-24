package main

import "fmt"

func main() {
	var max int

	if _, err := fmt.Scan(&max); err != nil {
		return
	}

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
