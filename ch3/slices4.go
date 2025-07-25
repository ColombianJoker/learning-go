package main

import "fmt"

func main() {
	x := make([]int, 5)
	x = append(x, 10)
	fmt.Println("\nx:", x, len(x), cap(x))

	y := make([]int, 5, 10)
	y = append(y, 10)
	fmt.Println("\ny:", y, len(y), cap(y))

	z := make([]int, 0, 10)
	z = append(x, 5, 6, 7, 8)
	fmt.Println("\nz:", z, len(z), cap(z))
}
