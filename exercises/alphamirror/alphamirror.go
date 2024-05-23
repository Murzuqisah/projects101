// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

/* The case of the letter remains unchanged, for example :

'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

The final result will be followed by a newline ('\n').

If the number of arguments is different from 1, the program prints a new line.
*/

package main

import (
	"fmt"
	"os"
)

func AlphaMirror(str string) string {
	s := ""
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			ch = 'z' - (ch - 'a')
			s += string(ch)
		} else if ch >= 'A' && ch <= 'Z' {
			ch = 'Z' - (ch - 'A')
			// str = append(str, )
			s += string(ch)
		} else {
			s += string(ch)
		}
	}
	return s
}

func main() {
	str := os.Args[1]
	result := AlphaMirror(str)
	fmt.Println(result)
}
