// Write a program that displays its first argument, if there is one.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := len(os.Args)

	if argument > 1 {
		var concatenated []rune
		for _, r := range os.Args[1] {
			concatenated = append(concatenated, r)
		}
		for _, r := range concatenated {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
