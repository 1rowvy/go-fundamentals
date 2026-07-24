package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string

	fmt.Scan(&x)

	if n, err := strconv.Atoi(x); err != nil {
		fmt.Print("bad")
	} else {
		fmt.Println("ok", n)
	}

}
