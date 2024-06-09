// Write a program that displays its first argument, if there is one.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	for _, char := range args[1] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// alternative II

// func main() {
// 	argument := len(os.Args)

// 	if argument > 1 {
// 		var concatenated []rune
// 		for _, r := range os.Args[1] {
// 			concatenated = append(concatenated, r)
// 		}
// 		for _, r := range concatenated {
// 			z01.PrintRune(r)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
