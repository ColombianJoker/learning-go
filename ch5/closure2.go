package main

import "fmt"

func makeMul(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMul(2)
	threeBase := makeMul(3)

	for i := 0; i <= 3; i++ {
		fmt.Println("i:", i, twoBase(i), threeBase(i))
	}
}
