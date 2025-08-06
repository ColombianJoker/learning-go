package main

import "fmt"

func main() {
	x := "hello"
	pointerToX := &x

	fmt.Println("x:", x)
	fmt.Println("pointerToX:", pointerToX)
	fmt.Println("*pointerToX:", *pointerToX)

	y := 10
	pointerToY := &y
	fmt.Println("*pointerToY:", *pointerToY)
	z := 5 + *pointerToY
	fmt.Println("z:", z)
}
