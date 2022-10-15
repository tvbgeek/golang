/*

2427. Number of Common Factors
Easy

177

1

Add to List

Share
Given two positive integers a and b, return the number of common factors of a and b.

An integer x is a common factor of a and b if x divides both a and b.

 

Example 1:

Input: a = 12, b = 6
Output: 4
Explanation: The common factors of 12 and 6 are 1, 2, 3, 6.
Example 2:

Input: a = 25, b = 30
Output: 2
Explanation: The common factors of 25 and 30 are 1, 5.
 

Constraints:

1 <= a, b <= 1000

*/

func commonFactors(a int, b int) int {
    num := 0
    var len1 int
    if a < b {
        len1 = a
    } else {
        len1 = b
    }
    
    i := 1
    for i <= len1 {
        if a % i == 0 && b % i == 0 {
            num++
        }
        i++
    }
    return num
}