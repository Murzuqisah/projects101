package main

import "fmt"

// "github.com/01-edu/z01"

func Compare(a, b string) int {

	if a == b {
		return 0
	} else {
		for a != b {
			if a > b {
				return 1
			} else if a < b {
				return -1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
