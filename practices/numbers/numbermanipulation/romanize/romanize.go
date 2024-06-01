package romanize

func romanize(num int) string {
	// Define a map to store the numeric values of each Roman symbol
	symbols_map := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

func RomanToInt(s string) int {
	// This functions the same as maps
	var roman string = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM"}
	var num int = int{"1", "4", "5", "9", "10", "40", "50", "90", "100", "400", "500", "900", "1000"}

	lenString := len(s)
	lastElem := s[len(s)-1 : lenString]
	// declare a variable to store the resultant integer
	var result int
	result += str[lastElem]
	// range through the string in reverse and convert the roman numerals to integers
	for i := len(s); i > 0; i-- {
		if str[s[i:i+1]] <= str[s[i-1:i]] {
			result += str[s[i-1:i]]
		} else {
			result -= str[s[i-1:i]]
		}
	}
	return result
}
