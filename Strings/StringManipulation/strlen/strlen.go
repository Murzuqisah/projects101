package main

import (
	"fmt"
)

func Strlen(str string) int {
	total := []rune(str)
	return len(total)
}

func main() {
	str := "Hello world"
	word := Strlen(str)
	fmt.Println(word)
}
