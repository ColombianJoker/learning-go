package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("'%s' is a short word!\n", word)
		case 5:
			wordLen := len(word)
			fmt.Printf("'%s' is exactly the right length: %d\n", word, wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("'%s' is a long word\n", word)
		}
	}
}
