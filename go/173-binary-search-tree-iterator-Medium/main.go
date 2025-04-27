/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    stack := []*TreeNode{}
    pushLeft(&stack, root)
    return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
    last := len(this.stack) - 1
    node := this.stack[last]
    this.stack[last] = nil // avoid memory leak
    this.stack = this.stack[:last]
    pushLeft(&this.stack, node.Right)
    return node.Val
}

func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0    
}

func pushLeft(stack *[]*TreeNode, node *TreeNode) {
    for node != nil {
        *stack = append(*stack, node)
        node = node.Left
    }
}