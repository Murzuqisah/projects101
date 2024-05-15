package main

import (
	"fmt"
	"numbers/prime"
	"numbers/minmax"
	"numbers/sumdiff"
	"numbers/twosums"
	"numbers/itoa"
)

func main() {

	// IsPrime()
	result := prime.IsPrime(7)
	fmt.Println(result)

	// MaxMin()
	max, min := minmax.MaxMin(7, 12)
	fmt.Println(max, min)

	// SumDiff()
	sum, diff := sumdiff.SumDiff(10, 5)
	fmt.Println("Sum = ",sum, "Difference = ", diff)

	// TwoSums
	//var result1 []int
	numbers := []int{789, 123, 345, 890, 567}
	target := int(912)
	result1 := twosums.TwoSums(numbers, target)
	fmt.Println(result1)

	// Itoa()
	v := int(23456543)
	result2 := itoa.Itoa(v)
	fmt.Println(result2)

}
