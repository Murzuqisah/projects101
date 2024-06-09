// Write a function Max that will return the maximum value in a slice of integers. If the slice is empty it will return 0.

package max

func Max(arr []int) int {
	var max int

	// initialize the first index as the maximum of the array
	max = arr[0]

	// loop through the array to find a maximum value greater than the first index, and return its value
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	return max
}
