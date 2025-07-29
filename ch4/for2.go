package main

import "fmt"

func main() {
	i := 1
	for i < 100 { // Condition-only for
		fmt.Println("i:", i)
		i = i * 2
	}
}
