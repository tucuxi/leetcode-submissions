/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    return build(inorder(root))
}

func inorder(root *TreeNode) []*TreeNode {
    var ordered []*TreeNode
    var traverse func(*TreeNode)
    
    traverse = func(root *TreeNode) {
        if root != nil {
            traverse(root.Left)
            ordered = append(ordered, root)
            traverse(root.Right)
        }
    }

    traverse(root)
    return ordered
}

func build(nodes []*TreeNode) *TreeNode {
    if len(nodes) == 0 {
        return nil
    }
    m := len(nodes) / 2
    node := nodes[m]
    node.Left = build(nodes[:m])
    node.Right = build(nodes[m + 1:])
    return node
}