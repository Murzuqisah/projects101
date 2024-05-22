package main

import (
	"github.com/01-edu/z01"
)

func main() {
	printAlphabet()
	z01.PrintRune('\n') // Print newline after printing the alphabet
}

func printAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a')%2 == 0 {
			z01.PrintRune(i) // Convert even letters to uppercase
		} else {
			z01.PrintRune(i - 32) // Keep odd letters in lowercase
		}
		// fmt.Printf("Debug: Printed rune %c\n", i)
	}
}