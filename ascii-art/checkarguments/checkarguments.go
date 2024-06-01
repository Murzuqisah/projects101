package checkarguments

import (
	"fmt"
	"os"
	"strings"
)

func CheckArguments(str string) string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . 'argument' asciiart_file.txt")
		os.Exit(0)
	}

	arguments := os.Args[1]
	argument := strings.Split(arguments, "\n")

	for _, arg := range argument {
		for _, chr := range arg {
			if chr < 32 || chr > 126 {
				fmt.Println("Error : Non ascii/printable characters found")
				os.Exit(0)
			}
		}
	}
	return str
}
