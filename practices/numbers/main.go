package main

import "os"

func main() {
	arg := os.Args[1]

	var roman string = ""
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 100}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for i, char := range arg {
		for i := 1; i <= 3900; i++ {
		}
	}
}
