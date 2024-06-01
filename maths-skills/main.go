package main

import (
	"fmt"
	"os"

	"mathskills/functions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	avg := functions.Average(nil)
	fmt.Println("Average:", avg)

	mid := functions.Median()
	fmt.Println("Median:", mid)

	variance, deviation := functions.StandardDeviation()
	fmt.Println("Variance:", variance, "\nStandard Deviation:", deviation)
}
