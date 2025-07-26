package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println("Orcas:", totalWins["Orcas"])
	fmt.Println("Lions:", totalWins["Lions"])
	totalWins["Kittens"]++
	fmt.Println("Kittens:", totalWins["Kittens"])
	totalWins["Lions"] += 2
	fmt.Println("totalWins:", totalWins)
}
