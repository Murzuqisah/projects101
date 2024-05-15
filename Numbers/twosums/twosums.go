package twosums

// Hash map method
func TwoSums(nums []int, target int) []int{
	if len(nums) <= 1 {
		return []int{-1, -1}
	}
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if k, l:= m[num]; l {
			return []int{k, i}
		}
		m[target-num] = i
	}
	return []int{-1, -1}
}
