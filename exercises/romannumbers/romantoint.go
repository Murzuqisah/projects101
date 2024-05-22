// Implement the From<u32> trait to create a roman number from a number. The roman number should be in subtractive notation (the common way to write roman number I, II, III, IV, V, VI, VII, VIII, IX, X ...).

// Start by defining the digits as RomanDigit (Nulla is 0).

// Next define RomanNumber as a wrapper to a vector of RomanDigit, and implement the Trait From<u32>.

package romannumbers

var romanInt = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func RomanNumber(roman string) int {

	// fmt.Println(len(roman))
	// constraints I <= num <= 3999
	strLen := len(roman)
	var LastElem = roman[len(roman)-1 : strLen]
	var result int

	result = know[LastElem]

	for i := len(roman) - 1; i > 0; i-- {
		if know[roman[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

// use roman_numbers::RomanNumber;
