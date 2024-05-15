package twosums

import (
	"fmt"
	"sort"
)

// sorting method
func TwoSumSorted(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{-1, -1}
	}

	// sort the nums
	sort.Ints(nums)
	fmt.Println(nums)

	for i, num := range nums {
		targetSum := target - num
		if k := sort.SearchInts(nums, targetSum); k < len(nums) && nums[k] == targetSum && k != i {
			return []int{i, k}
		}
	}
	return []int{-1, -1}
}