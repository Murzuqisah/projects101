package main

import "fmt"

func LastWord(s string) string {
	// word := string(s)[len(s)-1]

	var lastWord string

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			lastWord += string(s[i])
		} else {
			if i+1 < len(s) && s[i+1] != ' ' {
				lastWord = ""
			}
		}
	}
	return lastWord
}

func main() {
	str := " hello"
	result := LastWord(str)
	fmt.Println(result)
}
