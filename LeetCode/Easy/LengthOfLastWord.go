/*

58. Length of Last Word
Easy

1919

125

Add to List

Share
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.

 

Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.
 

Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.

*/

func lengthOfLastWord(s string) int {
    len1 := len(s)
    count := 0
    
    i := len1 - 1
    for i >= 0 && s[i] == ' ' {
        i--
    }
    
    for i >= 0 && s[i] != ' '{
        i--
        count++
    } 
    return count
}
