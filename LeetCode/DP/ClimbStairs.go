/*

Climbing Stairs

Solution
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

 

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 

Constraints:

1 <= n <= 45

*/

func climbStairs(n int) int {
    if n == 0 {
        return 0
    }    
    if n == 1 {
        return 1
    }
    
    arr := make([]int, 0)    
    // base condition
    arr = append(arr, 1)
    arr = append(arr, 2)
    
    i := 2
    for i < n {
        arr = append(arr, arr[i-1] + arr[i-2])
        i++
    }
    return arr[n-1]
}
