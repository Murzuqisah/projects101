package max

func Max(arr []int) int {
	var max int

	// initialize the first index as the maximum of the array
	max = arr[0]

	// loop through the array
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	return max
}