/* 

Remove Nth Node From End of List

Solution
Given the head of a linked list, remove the nth node from the end of the list and return its head.

 

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 

Follow up: Could you do this in one pass?

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 {
        return head
    }
    
    first := head
    second := head
    i := 0
    for i < n {
        second = second.Next
        i++
    }
    for second != nil && second.Next != nil {
        second = second.Next
        first = first.Next
    }
    
    if first.Next == nil {
        return nil    
    }
    
    fmt.Println(first.Val)
    fmt.Println(second)
    
    // we will hit this case when n is same as number of elements in LL
    if second == nil {
        return head.Next   
    } else {
        first.Next = first.Next.Next
    }
    
    return head
}
