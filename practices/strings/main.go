package main

import (
	"fmt"

	"spin/capitalizefirst"
)

func main() {
	// args := os.Args[1:]

	// for i, arg := range args {
	// 	for i := 0; i < len(arg); i++ {
	// 		if len(arg[i]) >= 5 {
	// 		}
	// 	}
	// }
	str := capitalizefirst.CapitalizeFirst("hello, how are you doing?")
	fmt.Println(str)
}
