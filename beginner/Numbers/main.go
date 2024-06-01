package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"numbers/calculator"
)

func main() {
	product := calculator.Product(32132.0, 42452.0)
	sum := calculator.Sum(123, 342)
	difference := calculator.Difference(9876, 43242)
	quotient := calculator.Quotient(45678, 5678)

	fmt.Println(product)
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(quotient)

	con := calculator.Convertor(3, "Blue")
	fmt.Println(con)

	user := "Sammy"
	lines := 50

	// You have to convert the variable lines to a string value before concatenating
	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	// converting from float to string value
	fmt.Println(fmt.Sprint(421.304))
	f := 59.53
	fmt.Println(fmt.Sprint(f))

	// concatenating the float with strings
	fmt.Println("Sammy has " + fmt.Sprint(f) + " points.")

	// converting strings to numbers
	lines_yesterday := "50"
	lines_today := "1080"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)

	// converting strings and bytes
	a := "my string"

	b := []byte(a)

	c := string(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
