package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]

	fmt.Println("s:", s)
	fmt.Println("b:", b)
	fmt.Printf("b: '%c'\n", b)
}
