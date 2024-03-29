/* 

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
 

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

*/

func twoSum(nums []int, target int) []int {
    len1 := len(nums)
    dc := make(map[int][]int, 0)
    //res := make([]int,0)
    res := make([]int,2)
    
    for i := 0; i < len1; i++ {
        dc[nums[i]] = append(dc[nums[i]], i)
    }
    
    //fmt.Println(dc)
    
    for i := 0; i < len1; i++ {
        // the following if condition is not needed for handling case
        //[-1,-2,-3,-4,-5] target -8
        //if nums[i] <= target {
            val, prs := dc[target - nums[i]]
            if prs {
                if nums[i] != target - nums[i] {
                    //res = append(res, i)
                    //res = append(res, val[0])
                    res[0] = i
                    res[1] = val[0]
                    return res
                }
            } 
            
            if prs && nums[i] == (target - nums[i]) && len(val) > 1 {
                //res = append(res, val[0])
                //res = append(res, val[1])
                res[0] = val[0]
                res[1] = val[1]
                return res
            }
        //}
    }  
    return []int{}
    
}
