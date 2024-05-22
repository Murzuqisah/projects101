package main

import (
	"fmt"
)

func FirstRune(n string) rune {
	return []rune(n)[0]
}

func main() {
	str := "hello world"
	result := FirstRune(str)
	fmt.Println(string(result))
}
