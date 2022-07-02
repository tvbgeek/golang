
/* Reverse Integer

Solution
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned). 

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
 

Constraints:

-2^31 <= x <= 2^31 - 1

*/

func reverse(x int) int {      
    
    nums := []int{}
    newnum := 0
    tens := 1
    negative := false
    
    if x < 0 {
        negative = true
        x = -x
    }
    
    //fmt.Printf("x %d", x)
    
    for x > 0 {
        nums = append(nums, x % 10)
        x = x/10
        tens = tens * 10
    } 
    
    tens = tens/10
    
    for _, val := range nums {
        newnum = newnum + val * tens
        tens = tens/10
    }
    
    // tricky stuff you need to check reversed integer for underflow or overflow
    // not the input integer
    
    if negative {
        if -newnum < math.MinInt32 {
            return 0
        } else {
            return -newnum
        }
    }
       
    if newnum > math.MaxInt32 {
        return 0
    } else {
        return newnum
    }
}
