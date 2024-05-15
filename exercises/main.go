package main

import (
	"fmt"
	"exercises/max"
	// "exercises/power"
	"exercise/firstrune"
)

func main() {
	arrInt := []int{12, 345, 453, 2455, 8923, 742, 4682}
	max := max.Max(arrInt)
	fmt.Println(max)

	// ispowerof2()
	//num := power.IsPowerOf2(75)

	// first rune
	s := firstrune.FirstRune("Hello")
	fmt.Println(s)

}