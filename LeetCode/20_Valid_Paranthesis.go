/*

20. Valid Parentheses
Easy

13969

652

Add to List

Share
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

*/


func isValid(s string) bool {
    len1 := len(s)
    
    if len1 % 2 != 0 {
        return false
    }
    
    arr := make([]rune, 0)

    for _, char := range s {
        //fmt.Println(arr)
        //fmt.Println(len(arr))
        if char == '(' || char == '{' || char == '[' {
            arr = append(arr, char)
        } else {     
            if char == ')' && (len(arr) == 0 || arr[len(arr)-1] != '(') {
                return false
            } else if char == '}' && (len(arr) == 0 || arr[len(arr)-1] != '{') {
                return false
            } else if char == ']' && (len(arr) == 0 || arr[len(arr)-1] != '[') {
                return false
            }
            arr = arr[:len(arr)-1]
          }
        }
    
    if len(arr) != 0 {
        return false
    } else {
        return true
    }
}
 
