package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println()
	fmt.Println("x:", x, "\tcap=", cap(x))
	y = append(y, 30)
	fmt.Println("x:", x, "\tcap=", cap(x))
	fmt.Println("y:", y, "\tcap=", cap(y))
}
