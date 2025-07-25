package main

import "fmt"

func main() {
	var x = []int{1, 5: 4, 5, 10: 100, 15}

	fmt.Println("x = ", x)
	x = append(x, 4)
	fmt.Println("x = ", x)

	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println("x = ", x)
}
