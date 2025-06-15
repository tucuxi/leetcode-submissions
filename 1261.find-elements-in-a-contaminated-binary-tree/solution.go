/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    values map[int]bool
}

func Constructor(root *TreeNode) FindElements {
    values := make(map[int]bool)
    
    var traverse func(*TreeNode, int)
    traverse = func(root *TreeNode, next int) {
        if root != nil {
            values[next] = true
            traverse(root.Left, 2 * next + 1)
            traverse(root.Right, 2 * next + 2)
        }
    }
    
    traverse(root, 0)    
    return FindElements{values}
}

func (this *FindElements) Find(target int) bool {
    return this.values[target]
}