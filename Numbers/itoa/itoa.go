package itoa

func Itoa(num int) string {
	var result []int
	lastDigit := 0

	// split the digits into int slices
	for num != 0{
		lastDigit = num%10
		result = append(result, lastDigit)
		num = num/10
	}

	// convert the int slices int a string
	var s string
	for i := len(result)-1; i >= 0; i-- {
		s += string(int32(result[i]) + '0')
	}
	return s	
}