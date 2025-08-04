package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, dividend % divisor, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("5 div 2:", result, remainder)
	result, remainder, err = divAndRemainder(5, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("5 div 0:", result, remainder)
}
