package main

// Write a program that displays the alphabet in reverse, with even
// letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

import "github.com/01-edu/z01"

func PrintRevAlpham() {
	for i := 'z'; i >= 'a'; i-- {
		if (i+'z')%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i - 32)
		}
	}
}

func main() {
	PrintRevAlpham()
	z01.PrintRune('\n')
}
