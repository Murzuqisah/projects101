//Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase,
// followed by a newline('\n').

package main

import "github.com/01-edu/z01"

func main(){
	// var result string
	
	var result []rune
	for i := 'a'; i <= 'z'; i++ {
		if (i - 'a')%2 == 0 { // Checking if the letter's index in the alphabet is even
			result += rune(i) - 32 // Convert even letters to uppercase
		} else {
			result += rune(i) // Keep odd letters in lowercase
		}
	}
	z01.PrintRune(i)
	z01.PrintRune('\n')

}

func ToLower(str rune) rune {
	var result rune

	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		}
	}
	return result

}



func ToUpper(str rune) rune {
	var result rune
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		}
	}
	return result
}
