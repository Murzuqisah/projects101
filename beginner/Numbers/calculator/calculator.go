package calculator

import (
	"fmt"
	"strconv"
)

func Product(num1, num2 int64) int64 {
	product := num1 * num2

	return product
}

func Sum(num1, num2 int64) int64 {
	sum := num1 + num2

	return sum
}

func Quotient(num1, num2 int64) int64 {
	quotient := num1 / num2

	return quotient
}

func Difference(num1, num2 int64) int64 {
	difference := num2 - num1

	return difference
}

// func Calculator(s string) int {
// 	args := os.Args[1:]
// 	arg := strings.Split(args[1:])

// 	//var operator string

// 	operator := []string{"+","-","/","*"}

// 	for i, char := range arg {
// 		if char[i] == operator {

// 		}
// 	}
// }

/* practice

#1 converting between int types

var index int8 = 15
var bigIndex int32
bigIndex = int32(index)
	fmt.Println(bigIndex)

output:
15

explanation:
	fmt.Printf("index data type: %T\n", index)
	fmt.Printf("bigIndex data type: %T\n", bigIndex)

output:
index data type: int8
bigIndex data type: int32

%T is the print verb for variable data type, not the actual value of the variable


#2 converting int to float

syntax:

var x int64 = 57
var y float64 = float64(x)
	fmt.Println("%.2f\n", y)

output:
57.00


#3 converting float to int

syntax:
var f float64 = 390.8
var i int = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

output:
f = 390.80
i = 390

#4 number converted through division

//example1
a := 5 / 2
	fmt.Println(a)

output:
	2

//example2
a := 5.0 / 2
	fmt.Println(a)

output:
	2.5

#5 Converting numbers to strings

a = strconv.Itoa(12)
	fmt.Printf("%q\n", a)

output:
	"12"


#6 COnverting float to string
Use fmt.Sprint() when you pass a float and a string value of the float will be returned


#7 converting strings and bytes
strings in go are stored as a slices of byte
You can convert between a slice of bytes and a string by wrapping it in the cooresponding conversions of
[]byte() and string


*/

func Convertor(x int, a string) int {
	x = 123
	var y float64 = float64(x)

	// using "%.2y" defines the number of decimal places that the variable y would have
	// fmt.Printf() formats the float with 2 dp
	// fmt.Printf("%.2y\n", y)

	// var a string
	// convert number to string
	a = strconv.Itoa(12)
	fmt.Printf("%q\n", a)

	/*output:
	"12"
	*/
	return int(y)
}
