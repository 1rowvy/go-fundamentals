package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var x1, y1, x2, y2 int

	fmt.Scan(&x1, &y1, &x2, &y2)

	p1 := Point{x1, y1}
	p2 := Point{x2, y2}

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	fmt.Print(dx*dx + dy*dy)
}
