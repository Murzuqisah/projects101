package twosums_test

import (
	"sort"
	"testing"
)

// func BenchmarkTwoSum(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		numbers.TwoSum(nums, target)
// 	}
// }

func BenchmarkTwoSumSorted(b *testing.B) {
	nums:= 
	sort.Ints(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numbers.TwoSumSorted(nums, target)
	}
}
