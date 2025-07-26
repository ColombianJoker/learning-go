package main

import "fmt"

func main() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	fmt.Println("x:", x, "\tcap(x):", cap(x))
	fmt.Println("y:", y, "\tcap(y):", cap(y))
	fmt.Println("z:", z, "\tcap(z):", cap(z))

	fmt.Println()
	y = append(y, 30, 40, 50)
	fmt.Println("x:", x, "\tcap(x):", cap(x))
	fmt.Println("y:", y, "\tcap(y):", cap(y))
	fmt.Println("z:", z, "\tcap(z):", cap(z))

	fmt.Println()
	// x = append(x, 60)
	fmt.Println("x:", x, "\tcap(x):", cap(x))
	fmt.Println("y:", y, "\tcap(y):", cap(y))
	fmt.Println("z:", z, "\tcap(z):", cap(z))

	fmt.Println()
	z = append(z, 70)
	fmt.Println("x:", x, "\tcap(x):", cap(x))
	fmt.Println("y:", y, "\tcap(y):", cap(y))
	fmt.Println("z:", z, "\tcap(z):", cap(z))
}
