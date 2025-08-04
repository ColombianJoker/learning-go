package main

import "fmt"

func main() {
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d²=%d\n", i, i*i)
	}
	fmt.Println()
}
