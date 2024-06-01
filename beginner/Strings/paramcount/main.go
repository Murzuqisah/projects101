package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	paramCount := len(args)
	var result string

	// fmt.Println(paramCount)

	result = append(result, string(int32(paramCount)+'0'))

	z01.PrintRune(result)

	// for multiple parameters, printing in rune
}

// func Itoa(num int) string {

// }