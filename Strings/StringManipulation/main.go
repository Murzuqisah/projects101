package main

import "fmt"

func main() {
	sentence := "Hello world"
	counter := 0

	for index, letter := range sentence {
		counter++
		fmt.Printf("index: %v letter: %c\n", index, letter)
	}
	fmt.Printf("counter value: %v\n", counter)
	fmt.Println()

	// #2 []byte

	word := []byte("Hello")
	count:= 0
	for index, letter := range word {
		count++
		fmt.Printf("index: %v letter: %c\n", index, letter)
	}
	fmt.Printf("count: %v\n", count)
	fmt.Println()


	words := []rune("Hello")
	countr := 0

	for i, ch := range words {
		countr++
		fmt.Printf("i: %v ch: %c\n", i, ch)
	}
}