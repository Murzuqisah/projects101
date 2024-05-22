package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := len(os.Args)

	if argument > 1 {
		var concatenated []rune
		// runes := []rune(os.Args[argument-1])
		for _, r := range os.Args[argument-1] {
			concatenated = append(concatenated, r)
		}
		for _, r := range concatenated {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

/*
# ii. alternative 2

package main

import (
	"os"
	"fmt"
)

func main() {
	argument := len(os.Args)
	if argument > 1 {
		fmt.Println(os.Args[argument-1])
	}
}

# iii. alternative 3

package main

import (
	"os"
	"github.com/01-edu/z01"
)

// func main() {
// 	argument := len(os.Args)

// 	for i := argument - 1; i > 0; i-- {
// 		args := os.Args[i]
// 		for _, char := range args {
// 			z01.PrintRune(char)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

*/
