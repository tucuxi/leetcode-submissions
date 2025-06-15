/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    stack := []*TreeNode{}
    for p := root; p != nil; p = p.Right {
        if p.Right != nil {
            stack = append(stack, p.Right)
        }
        if p.Left != nil {
            p.Left, p.Right = nil, p.Left
            continue
        }
        if len(stack) > 0 {
            p.Right = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
    }
}