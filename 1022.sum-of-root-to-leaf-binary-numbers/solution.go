/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0, 0)
}

func dfs(root *TreeNode, acc int, sum int) int {
    if root == nil {
        return sum
    }
    newAcc := 2 * acc + root.Val
    if root.Left == nil && root.Right == nil {
        return sum + newAcc
    }
    return dfs(root.Left, newAcc, sum) + dfs(root.Right, newAcc, sum)
}

