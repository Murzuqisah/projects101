package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func StandardDeviation() (int, int) {
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
	var sum, mean, deviation, variance float64

	// read the file line by line
	for scanner.Scan() {
		// convert the strings to integer values
		value := scanner.Text()
		number, err := strconv.ParseFloat(value, 64)
		CheckNilError(err, "Error converting string to floating point number")

		digits = append(digits, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	// Average = (sum/size of array)
	sum = 0
	total := float64(len(digits))

	for _, digit := range digits {
		sum += digit
	}

	// calculate the mean value
	mean = sum / total

	// Calculate the squared difference from the mean
	// range over the array values to calculate the sd variable
	for _, values := range digits {
		deviation += math.Pow(values-mean, 2)
	}

	// calculate the variance
	deviation = deviation / total

	// calculate the standard deviation
	variance = math.Sqrt(deviation)

	return int(math.Round(deviation)), int(math.Round(variance))
}
