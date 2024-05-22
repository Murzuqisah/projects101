package main

import "fmt"

func LastRune(str string) rune {
	return []rune(str)[len(str)-1]
}

func main() {
	str := "hello world"
	result := LastRune(str)
	fmt.Println(string(result))
}
