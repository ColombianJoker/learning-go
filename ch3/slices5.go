package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x, len(x), cap(x))
	fmt.Println("y:", y, len(y), cap(y))
	fmt.Println("z:", z, len(z), cap(z))
	fmt.Println("d:", d, len(d), cap(d))
	fmt.Println("e:", e, len(e), cap(e))
}
