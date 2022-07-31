/*

Maximum Subarray

Solution
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

 

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
 

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

*/


func maxSubArray(nums []int) int {
    len1 := len(nums)
    
    if len1 == 1 {
        return nums[0]
    }
    
    allNegative := true
    // handle the all negative elements case separately
    i := 0
    maxElement := math.MinInt32
    for i < len1 {
        if nums[i] >= 0 {
            allNegative = false
            break
        }
        if nums[i] > maxElement {
            maxElement = nums[i]
        }
        i++
    }
    
    if allNegative {
        return maxElement
    }
    
    // regular flow
    maxSum := 0
    currSum := 0
    i = 0
    
    for i < len1 {
        currSum = 0
        for i < len1 {
            currSum = currSum + nums[i]
            if currSum >= 0 {
                if currSum > maxSum {
                    maxSum = currSum
                }
            } else {
                i++
                break
            }
            i++
        }
    }
    return maxSum
}
