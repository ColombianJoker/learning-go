package main

import "fmt"

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	ok, err := doubleEven(2)
	// fmt.Printf("ok: %d\terr:%d\n", ok, err)
	if err == nil {
		fmt.Println("doubleEven(2):", ok)
	} else {
		fmt.Println("doubleEven(2)->", err)
	}

	ok, err = doubleEven(3)
	// fmt.Printf("ok: %d\terr:%d\n", ok, err)
	if err == nil {
		fmt.Println("doubleEven(3):", ok)
	} else {
		fmt.Println("doubleEven(3)->", err)
	}
}
