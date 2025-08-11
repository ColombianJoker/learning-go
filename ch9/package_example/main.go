package main

import (
	"fmt"

	print "package_example/formatter"
	math "package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
