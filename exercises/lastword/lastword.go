package lastword

import (
	"os"
)

func LastWord(args string) string {
	args = os.Args[1]
	// word := string(s)[len(s)-1]
	var lastWord string

	for i := 0; i < len(args); i++ {
		if args[i] != ' ' {
			lastWord += string(args[i])
		} else {
			if i+1 < len(args) && args[i+1] != ' ' {
				lastWord = ""
			}
		}
	}
	return lastWord
}

// ver 2.0

// package lastword

// import (
// 	"os"
// 	"strings"
// )

// func LastWord(args string) string {
// 	words := strings.Fields(os.Args[1])
// 	if len(words) == 0 {
// 		return ""
// 	}
// 	return words[len(words)-1]
// }
