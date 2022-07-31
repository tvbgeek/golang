/*

House Robber

Solution
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

 

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
 

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400

*/

func getMax(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}

func rob(nums []int) int {
    len1 := len(nums)
    
    // corner case only one element or 2 elements in input array
    if len1 == 1 {
        return nums[0]
    } else if (len1 == 2) {
        return getMax(nums[0], nums[1])
    }
    
    arr := make([]int, len1)
    
    arr[0] = nums[0]
    arr[1] = getMax(nums[0], nums[1])
    
    i := 2    
    for i < len1 {
        arr[i] = getMax((nums[i] + arr[i-2]), arr[i-1])
        i++
    }
    return arr[len1-1]
}
