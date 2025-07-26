package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println("x:", x, "\tlen(x):", len(x), "\tcap(x):", cap(x))
	num := copy(x[:3], x[1:])
	fmt.Println("x:", x, "\tlen(x):", len(x), "\tcap(x):", cap(x))
	fmt.Println("num:", num)
}
