/*

First Unique Character in a String

Solution
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
 
Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1
 

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.

*/

func firstUniqChar(s string) int {
    
    // we need to keep track of index and count for each character
    dict := make(map[rune][2]int, len(s))
        
    for i, val := range s {
        arr, prs := dict[val]
        if prs {
            arr[1] = arr[1] + 1
        } else {
            arr[0] = i
            arr[1] = 1
        }
        dict[val] = arr
    }
    
    // check character with min_index and repetition == 1 
    min_index := math.MaxInt32
    for _, v := range dict {
        if v[1] == 1 && v[0] < min_index {
            min_index = v[0]
        }
    }
    
    if min_index == math.MaxInt32 {
        return -1
    } else {
        return min_index
    }    
}
