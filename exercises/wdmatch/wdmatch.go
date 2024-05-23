// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string.
// This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	WdMatch("quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoix")
}

func WdMatch(str1, str2 string) {
	if IsMatch(str1, str2) {
		for _, v := range str1 {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func IsMatch(str1, str2 string) bool {
	// initialize character matches
	matches := 0
	// loop through str2 checking the matches in every index, in order.
	for i := 0; i < len(str2); i++ {
		if matches < len(str1) {
			if str1[matches] == str2[i] {
				matches++
			}
		}
	}

	return len(str1) == matches
}
