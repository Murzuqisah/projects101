package main

import (
	"exercise/firstrune"
	"fmt"

	"exercises/lastword"
	"exercises/max"
	"exercises/power"
)

func main() {
	arrInt := []int{12, 345, 453, 2455, 8923, 742, 4682}
	max := max.Max(arrInt)
	fmt.Println(max)

	num := power.IsPowerOf2(75)
	fmt.Println(num)

	// first rune
	s := firstrune.FirstRune("Hello")
	fmt.Println(s)

	// last word

	str := " hello"
	result := lastword.LastWord(str)
	fmt.Println(result)
}
