/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type triplet struct{val, left, right int}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    ids := map[triplet]int{}
    nextId := 1
    count := map[int]int{}

    var traverse func(*TreeNode) int
    traverse = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        t := triplet{root.Val, traverse(root.Left), traverse(root.Right)}
        if id, exists := ids[t]; exists {
            if count[id] == 1 {
                res = append(res, root)
            }
        } else {
            ids[t] = nextId
            nextId++
        }
        count[ids[t]]++
        return ids[t]
    }

    traverse(root)
    return res
}