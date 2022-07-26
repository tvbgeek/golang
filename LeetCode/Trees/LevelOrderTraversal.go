/*

Binary Tree Level Order Traversal

Solution
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

 

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    curr := make([]*TreeNode, 0)
    next := make([]*TreeNode, 0)
    
    curr = append(curr, root)
    rows := 0
    cols := 0
    res := [][]int{}
    
    for len(curr) > 0 {
        res = append(res, []int{})
        fmt.Println(res)
        for _, node := range curr {
            //fmt.Printf("node %v rows %d cols %d", node, rows, cols)
            res[rows] = append(res[rows], node.Val)
       
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)   
            }        
        }
        //fmt.Println("next:", next)
        curr = next
        next = nil
        //fmt.Println("curr:", curr)
        rows++    
    }
    return res
}
