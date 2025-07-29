package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Outside if {} x:", x)
	if x > 5 {
		x, y := 5, 20
		fmt.Println("Inside if {} x, y:", x, y)
		x = 15
		fmt.Println("Inside if {} x, y:", x, y)
	}
	fmt.Println("Outside if {} x:", x)
}
