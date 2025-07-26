package main

import "fmt"

func main() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)

	fmt.Println("a:", a)
	fmt.Printf("a: '%c'\n", a)
	fmt.Println("s:", s)
	fmt.Println("b:", b)
	fmt.Printf("b: '%c'\n", b)
	fmt.Println("s2:", s2)
}
