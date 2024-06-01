package main

import "fmt"

func IntToRomans(number int) string {
	var nums = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1
	var roman = ""

	for number > 0 {
		for nums[index] <= number {
			roman += romans[index]
			number -= nums[index]
		}
		index -= 1
	}
	return roman
}

func main() {
	number := 14
	result := IntToRomans(number)
	fmt.Println(result)
}
