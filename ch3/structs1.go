package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	bob := person{}
	julia := person{
		"Julia", 40, "cat",
	}
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println("bob:", bob)
	fmt.Println("julia:", julia)
	fmt.Println("beth:", beth)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println("pet:", pet)
	fmt.Println("pet.name:", pet.name)
	fmt.Println("pet.kind:", pet.kind)
}
