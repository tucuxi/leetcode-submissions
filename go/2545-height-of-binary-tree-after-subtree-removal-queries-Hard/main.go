/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func treeQueries(root *TreeNode, queries []int) []int {
    heights := makeHeights(root)
    results := make(map[int]int)

    var dfs func(*TreeNode, int, int)

    dfs = func(node *TreeNode, depth, maxValue int) {
        if node != nil {
            results[node.Val] = maxValue
            dfs(node.Left, depth + 1, max(maxValue, depth + 1 + heights[node.Right]))
            dfs(node.Right, depth + 1, max(maxValue, depth + 1 + heights[node.Left]))
        }
    }

    dfs(root, 0, 0)

    res := make([]int, len(queries))
    for i, q := range queries {
        res[i] = results[q]
    }
    return res
}

func makeHeights(root *TreeNode) map[*TreeNode]int {
    heights := make(map[*TreeNode]int)
    heights[nil] = -1

    var dfs func(*TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return -1
        }
        res := max(dfs(node.Left), dfs(node.Right)) + 1
        heights[node] = res
        return res
    }

    dfs(root)
    return heights
}