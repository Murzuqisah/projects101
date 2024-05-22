package main

import (
	"fmt"
)

func main() {
	printAlphabet()
}

func printAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a')%2 == 0 {
			fmt.Printf("%c", i-32) // Convert even letters to uppercase
		} else {
			fmt.Printf("%c", i) // Keep odd letters in lowercase
		}
	}
	fmt.Println() // Print newline after printing the alphabet
}

// ver 2.0

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	alphabet := generateAlphabet()
// 	displayAlphabet(alphabet)
// }

// func generateAlphabet() string {
// 	var result string
// 	for i := 'a'; i <= 'z'; i++ {
// 		if (i - 'a')%2 == 0 { // Checking if the letter's index in the alphabet is even
// 			result += string(i - 32) // Convert even letters to uppercase
// 		} else {
// 			result += string(i) // Keep odd letters in lowercase
// 		}
// 	}
// 	return result
// }

// func displayAlphabet(alphabet string) {
// 	fmt.Println(alphabet)
// }
