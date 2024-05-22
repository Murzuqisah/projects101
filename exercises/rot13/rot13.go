package main

import "fmt"

func Rot13(s string) string {
	result := ""
	for _, char := range s {
		if 'a' <= char && char <= 'z' {
			result += string((char-'a'+13)%26 + 'a')
		} else if 'A' <= char && char <= 'Z' {
			result += string((char-'A'+13)%26 + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	str := "z"
	result := Rot13(str)
	fmt.Println(result)
}
