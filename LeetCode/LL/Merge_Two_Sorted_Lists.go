/*

Merge Two Sorted Lists

Solution
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

 

Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
 

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    
    if list2 == nil {
        return list1
    }
    
    l1 := list1
    l2 := list2
    var res *ListNode
    
    if l1.Val <= l2.Val {
        res = l1
        l1 = l1.Next
    } else {
        res = l2
        l2 = l2.Next
    }
    
    curr := res
    
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            curr = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            curr = l2
            l2 = l2.Next
        }
    }
    
    if l1 == nil && l2 != nil {
        for l2 != nil {
            curr.Next = l2
            curr = l2
            l2 = l2.Next
        }                
    } 
    
    if l1 != nil && l2 == nil {
        for l1 != nil {
            curr.Next = l1
            curr = l1
            l1 = l1.Next            
        }
    }
    
    return res
}
