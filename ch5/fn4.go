package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(dividend, divisor int) (result int, remainder int, err error) {
	if divisor == 0 {
		return result, remainder, errors.New("cannot divide by zero")
	}
	result, remainder = dividend/divisor, dividend%divisor
	return result, remainder, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("5 ÷ 2:", result, remainder)
	result, remainder, err = divAndRemainder(5, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("5 div 0:", result, remainder)
}
