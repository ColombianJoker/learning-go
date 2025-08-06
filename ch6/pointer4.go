package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func stringp(s string) *string {
	return &s
}

func main() {
	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Petterson",
	}

	fmt.Println("p:", p)
}
