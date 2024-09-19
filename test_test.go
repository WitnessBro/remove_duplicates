package remove_duplicates

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{0, 0, 0, 1, 2}   // Input array 	// Value to remove
	var expectedNums = []int{0, 1, 2} // The expected answer with correct length.
	var k = removeDuplicates(nums)    // Calls your implementation

	assert.Equal(t, k, len(expectedNums))
	for i := 0; i < k; i++ {
		assert.Equal(t, nums[i], expectedNums[i])
	}
	fmt.Println(nums)
}
func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
