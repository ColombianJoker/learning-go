package main

import "fmt"

func main() {
	x := make([]int, 5)
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
}
