/*

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?

*/

package leetcode

func moveZeroes(nums []int) {
	len1 := len(nums)
	if len1 <= 1 {
		return
	}

	i := 0
	j := 0

	for i < len1 && j < len1 {
		for i < len1 && nums[i] != 0 {
			i++
		}
		for j < len1 && nums[j] == 0 {
			j++
		}
		if i < len1 && j < len1 && i < j {
			nums[i] = nums[j]
			nums[j] = 0
			i++
			j++
		} else {
			j++
		}
	}
}
