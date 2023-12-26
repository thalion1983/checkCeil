package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int = 6

	for a = 12; a <= 18; a++ {
		fmt.Printf("Division: %d/%d\n===============\n", a, b)
		noConv  := a / b                                   // No conversion at all
		oneConv := int(math.Ceil(float64(a / b)))          // Convert the result of division to float64
		twoConv := int(math.Ceil(float64(a) / float64(b))) // Convert a and b to float64 before division
		fmt.Printf("              No float conversion: %d\n", noConv)
		fmt.Printf("        Division result converted: %d\n", oneConv)
		fmt.Printf("a and b converted before division: %d\n\n", twoConv)
	}
}
