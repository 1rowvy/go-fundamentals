package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	res, err := safeDivide(a, b)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("result: %d", res)
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}

	return a / b, nil
}
