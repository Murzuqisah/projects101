package main

import "fmt"

func main() {
	fmt.Println(ReverseBits(192))
}

func ReverseBits(oct byte) byte {
	var count byte = 1
	var i byte
	var result byte = 0

	for i = 128; oct > 0; i = i / 2 {
		if i <= oct {
			result += count
			oct -= i
		}
		count *= 2
	}
	return result
}
