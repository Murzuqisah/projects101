//Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase,
// followed by a newline('\n').

package main

import "github.com/01-edu/z01"

func main(){
	// var result string
	
	for i := 'a'; i <= 'z'; i++ {

		if i % 2==0{
			char :=ToUpper(i)
			z01.PrintRune(char)
		} else if i % 2 == 1 {
			cha := ToLower(i)
			z01.PrintRune(cha)
		}
		// i = append(i, result)

	}
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