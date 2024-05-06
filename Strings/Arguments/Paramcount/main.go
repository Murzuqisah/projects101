package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := len(os.Args[1:])

	num := Itoa(args)
	Print(num)

}

func Itoa(num int) string {
	var result []int
	var mod int
	var div int
	div = num

	for div > 0 {
		mod = div % 10
		div = div / 10
		result = append(result, mod)
	}

	var char rune
	var numbers string
	var word string

	for i := len(result) - 1; i >= 0; i-- {
		char = int32(result[i]) + '0'
		numbers = string(char)
		word += numbers
	}
	return word

}

func Print(word string) {

	for _, s := range word {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
