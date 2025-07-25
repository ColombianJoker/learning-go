package main

import "fmt"

func main() {
	var x = []int{1, 5: 4, 5, 10: 100, 15}

	fmt.Println("Values in x:")
	for i, v := range x {
		fmt.Printf("x[%d] = %d\n", i, v)
	}

	fmt.Println("")

	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %d\n", i, x[i])
	}
}
