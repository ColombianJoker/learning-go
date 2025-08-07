package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Println("Uncategorized:", Uncategorized)
	fmt.Println("Personal:", Personal)
	fmt.Println("Spam:", Spam)
	fmt.Println("Social:", Social)
	fmt.Println("Advertisement:", Advertisements)
}
