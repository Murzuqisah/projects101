package main

import (
	"fmt"
)

func ActiveBits(n int) int {
	count := 0

	for i := 128; n > 0; i= i/2 {
		if i <= n {
			count++
			n -= i
		}
	}
	return count
}

func main() {
	fmt.Println(ActiveBits(163))
}
