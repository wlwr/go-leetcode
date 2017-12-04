package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []int{2, 7, 11, 15}
	target := 22
	result := []int{1, 3}
	if res := twoSum(nums, target); res[0] != result[0] || res[1] != result[1] {
		t.Fatalf("%v\n", res)
	}
}
