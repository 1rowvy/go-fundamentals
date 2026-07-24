package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := scanner.Text()

	strUp := strings.ToUpper(str)

	fmt.Println(strUp)
}
