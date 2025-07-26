package main

import "fmt"

func main() {
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raúl", "Ze"},
	}
	fmt.Println("teams:", teams, "\tlen(teams):", len(teams))

	tenMap := make(map[int][]string, 10)
	fmt.Println("tenMap:", tenMap, "\tlen(tenMap):", len(tenMap))
}
