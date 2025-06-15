/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {    
    return traverse(root, []byte{})
}

func traverse(root *TreeNode, path []byte) string {
    path = append([]byte{byte(root.Val) + 'a'}, path...)
    if root.Left == nil && root.Right == nil {
        return string(path)
    }
    var res string
    if root.Left != nil {
        res = traverse(root.Left, path)
    }
    if root.Right != nil {
        if r := traverse(root.Right, path); res == "" || r < res {
            res = r
        }
    }
    return res
}