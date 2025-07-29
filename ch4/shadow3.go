package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println("x:", x)
	fmt := "oops"
	fmt.Println("x:", x)
	// go run ch4/shadow3.go
	// # command-line-arguments
	// ch4/shadow3.go:11:6: fmt.Println undefined (type string has no field or method Println)
}
