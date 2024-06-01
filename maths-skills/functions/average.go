package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Average(array []float64) int {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		os.Exit(1)
	}

	fileName := os.Args[1]
	// read from data.txt file
	file, err := os.Open(fileName)
	CheckNilError(err, "Failed to open file")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// initialize array to store the elements
	var digits []float64

	// read the file line by line
	for scanner.Scan() {
		value := scanner.Text()
		// convert the strings to integer values
		number, err := strconv.ParseFloat(value, 64)
		CheckNilError(err, "Error converting string to integer")

		digits = append(digits, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error convering to integer")
		os.Exit(1)
	}

	// Average = (sum/size of array)
	sum := 0.0
	total := float64(len(digits))

	for _, digit := range digits {
		sum += digit
	}

	average := (sum / total)

	return int(math.Round(average))
}