package twosums_test

import (
	"testing"
	"sort"
	"numbers/twosums"
)


// func BenchmarkTwoSum(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		numbers.TwoSum(nums, target)
// 	}
// }

func BenchmarkTwoSumSorted(b *testing.B) {
	sort.Ints(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numbers.TwoSumSorted(nums, target)
	} 
}
