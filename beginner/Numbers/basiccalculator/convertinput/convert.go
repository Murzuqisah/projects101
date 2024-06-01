package convertinput

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertInputToType(input string) float64 {
	args := os.Args[1:]

	// iterate through the argument
	for i, arg := range args {
		strings.Split(arg, " ")
		if arg[i] == string() {
			arg = strconv.ParseFloat(string, float64)
			fmt.Printf("%")
		}
	}

	// convert the input from string to float64
}
