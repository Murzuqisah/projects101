// Write a function that counts the characters of a string and that returns that count.

package main

import "fmt"

func Strlen(str string) int {
	count := 0
	for i:= 0; i < len(str); i++ {
		count++
	}
	return count
}

func main() {
	str := "Hello World!"
	nb := Strlen(str)
	fmt.Println(nb)
}