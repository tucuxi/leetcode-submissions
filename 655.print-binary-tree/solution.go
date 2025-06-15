/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type queueEntry struct {
    node *TreeNode
    row, col int
}

func printTree(root *TreeNode) [][]string {
    m := height(root) + 1
    n := ipow(2, m) - 1
    res := make([][]string, m)
    for i := range m {
        res[i] = make([]string, n)
    }
    for q := []queueEntry{queueEntry{root, 0, (n - 1) / 2}}; len(q) > 0; q = q[1:] {
        node, r, c := q[0].node, q[0].row, q[0].col
        res[r][c] = strconv.Itoa(node.Val)
        offset := ipow(2, m - r - 2)
        if node.Left != nil {
            q = append(q, queueEntry{node.Left, r + 1, c - offset})
        }
        if node.Right != nil {
            q = append(q, queueEntry{node.Right, r + 1, c + offset})
        }
    }
    return res
}

func height(root *TreeNode) int {
    if root == nil {
        return -1
    }
    return max(height(root.Left), height(root.Right)) + 1
}

func ipow(b, e int) int {
    res := 1
    for range e {
        res *= b
    }
    return res
}