package main

import "fmt"

func main() {

	words := make(map[string]bool)

	for {
		var str string

		if _, err := fmt.Scan(&str); err != nil {
			break
		}

		words[str] = true
	}

	fmt.Print(len(words))

}
