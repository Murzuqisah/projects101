package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// create a struct to sort values
// type Digits struct {
// 	digits []int
// 	value int
// }

func Median() int {
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
		// convert the strings to integer values
		currentLine := scanner.Text()
		number, err := strconv.ParseFloat(currentLine, 64)
		CheckNilError(err, "Error converting string to floating point number")

		digits = append(digits, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error converting to float value")
		os.Exit(1)
	}
	// sort the data in ascending order
	sort.Float64s(digits)

	length := len(digits)
	if length == 0 {
		return 0
	}

	// for odd number of values, default output
	// if (length % 2) == 0
	// for even number of values
	// fmt.Println(int(math.Round((digits[(length/2)-1] + digits[length/2]) / 2)))
	if (length % 2) == 1 {
		return (int(math.Round((digits[(length/2)-1] + digits[length/2]) / 2)))
	}

	return int(math.Round(digits[length/2]))
}
