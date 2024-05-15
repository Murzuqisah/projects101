package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// read the command-line arguments after the program name
	args := len(os.Args[1:])

	// call the itoa function to convert the integer and then print out the num(as a string)
	num := Itoa(args)
	Print(num)

}

// converts int to string for undefined length of integers
func Itoa(num int) string {
	// declare variables to store the result, dividend, modulus
	var result []int
	var mod int
	var div int
	div = num

	// create a while loop to perform the calculations
	for div > 0 {
		mod = div % 10
		div = div / 10

		// append the modulus to result
		result = append(result, mod)
	}

	// declare variables to store the converted data types
	var char rune
	var numbers string
	var word string

	// convert int to rune and then to string
	for i := len(result) - 1; i >= 0; i-- {
		char = int32(result[i]) + '0'
		numbers = string(char)
		word += numbers
	}
	// return the string
	return word

}

// printing function
func Print(word string) {

	for _, s := range word {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
