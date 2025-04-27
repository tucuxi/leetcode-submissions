/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    var res []int

    t1, t2 := inorder(root1), inorder(root2)
    i1, i2 := 0, 0
    for i1 < len(t1) || i2 < len(t2) {
        if i2 == len(t2) || i1 < len(t1) && t1[i1] < t2[i2] {
            res = append(res, t1[i1])
            i1++
        } else {
            res = append(res, t2[i2])
            i2++
        }
    }
    return res
}

func inorder(root *TreeNode) []int {
    var res []int
    var traverse func(*TreeNode)
    
    traverse = func(root *TreeNode) {
        if root != nil {
            traverse(root.Left)
            res = append(res, root.Val)
            traverse(root.Right)
        }
    }
    traverse(root)
    return res
}