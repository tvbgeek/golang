/*

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.

*/

func isPalindrome(s string) bool {
    
    if len(s) <= 1 {
        return true
    }
    
    arr := make([]rune, 0)
    
    for _, char := range strings.ToLower(s) {
        if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
            arr = append(arr, char)
        }
    }
    
    fmt.Println(arr)
    
    len1 := len(arr)
    if len1 <= 1 {
        return true
    }
    
    for i, j := 0, len1 - 1; i <= j; i, j = i + 1, j - 1 {
        if arr[i] != arr[j] {
            return false
        }
    } 
    
    return true
}
