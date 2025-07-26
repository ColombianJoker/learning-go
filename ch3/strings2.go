package main

import "fmt"

func main() {
	var s string = "Hello 😀"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]

	fmt.Printf("s: '%s'\tlen(s): %d\n", s, len(s))
	fmt.Printf("s2: '%s'\tlen(s2): %d\n", s2, len(s2))
	fmt.Printf("s3: '%s'\tlen(s3): %d\n", s3, len(s3))
	fmt.Printf("s4: '%s'\tlen(s4): %d\n", s4, len(s4))
}
