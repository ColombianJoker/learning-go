package main

import (
	"fmt"
	"os"
	"strconv"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as an argument.")
		return
	}

	arg := os.Args[1]
	number, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number.\n", arg)
		return
	}

	for i := range countTo(number) {
		fmt.Println(i)
	}
}
