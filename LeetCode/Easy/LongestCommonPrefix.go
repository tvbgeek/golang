/*

14. Longest Common Prefix
Easy

10754

3447

Add to List

Share
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.

*/

func commonPrefix(s1, s2 string) string {
    r1 := []rune(s1)
    r2 := []rune(s2)
    
    //fmt.Println(r1, r2)
    
    var len1 int
    if len(r1) < len(r2) {
        len1 = len(r1)
    } else {
        len1 = len(r2)
    }
    
    c := ""
    i := 0
    for i < len1 {
        if r1[i] == r2[i] {
            //fmt.Println(c, r1[i], r2[i])
            c = c + string(r1[i])
        } else {
            break
        }
        i++
    }
    return c
}


func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 0 {
        return ""
    } else if len(strs) == 1 {
        return strs[0]
    }
    
    lcp := commonPrefix(strs[0], strs[1])
    //fmt.Println(lcp)
    len1 := len(strs)
    i := 2
    
    for i < len1 {
        lcp = commonPrefix(lcp, strs[i]) 
        i++
    }
    return lcp
}
