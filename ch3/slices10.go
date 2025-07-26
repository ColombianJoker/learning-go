package main

import "fmt"

func main() {
	x := []int{5, 6, 7, 8}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println("y:", y, "\tlen(y):", len(y), "\tcap(y):", cap(y))
	fmt.Println("num:", num, "\tlen(num):", len(num), "\tcap(num):", cap(num))
}
