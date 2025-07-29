package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz ")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz ")
			continue
		}
		fmt.Print(i, " ")
	}
}
