package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println("len(vals):", len(vals), "\tlen(intSet):", len(intSet))
	fmt.Println("intSet[5]:", intSet[5])
	fmt.Println("intSet[500]:", intSet[500])
	if intSet[100] {
		fmt.Println("100 is in intSet")
	}
	fmt.Println()
	fmt.Println("vals:", vals)
	fmt.Println("intSet:", intSet)
}
