/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode)
    noPredecessors := make(map[*TreeNode]bool)
    for _, d := range descriptions {
        parent, child, isLeft := d[0], d[1], d[2]
        p := nodes[parent]
        if p == nil {
            p = &TreeNode{Val: parent}
            nodes[parent] = p
            noPredecessors[p] = true
        }
        c := nodes[child]
        if c == nil {
            c = &TreeNode{Val: child}
            nodes[child] = c
        } else {
            delete(noPredecessors, c)
        }
        if isLeft == 1 {
            p.Left = c
        } else {
            p.Right = c
        }
    }
    for root := range noPredecessors {
        return root
    }
    return nil
}