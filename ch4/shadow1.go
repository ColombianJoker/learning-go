package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Outside if {} x:", x)
	if x > 5 {
		fmt.Println("Inside if {} x:", x)
		x := 5
		fmt.Println("Inside if {} x:", x)
	}
	fmt.Println("Outside if {} x:", x)
}
