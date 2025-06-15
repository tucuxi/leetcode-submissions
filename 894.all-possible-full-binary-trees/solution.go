/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    trees := make([][]*TreeNode, n + 1)
    trees[1] = append(trees[1], &TreeNode{})
    for i := 3; i <= n; i += 2 {
        for j := 1; j < i; j += 2 {
            k := i - j - 1
            for _, l := range trees[j] {
                for _, r := range trees[k] {
                    trees[i] = append(trees[i], &TreeNode{0, l, r})
                }
            }
        }
    }
    return trees[n]
}
