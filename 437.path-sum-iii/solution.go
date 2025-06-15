/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    return traverse(root, targetSum, []int{})
}

func traverse(root *TreeNode, targetSum int, path []int) int {
    if root == nil {
        return 0
    }
    count, sum := 0, 0
    path = append(path, root.Val)
    for i := len(path) - 1; i >= 0; i-- {
        sum += path[i]
        if sum == targetSum {
            count++
        }
    }
    return count + traverse(root.Left, targetSum, path) + traverse(root.Right, targetSum, path)
}