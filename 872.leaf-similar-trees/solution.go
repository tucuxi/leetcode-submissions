/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go leaves(root1, ch1)
    go leaves(root2, ch2)
    
    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2
        if ok1 != ok2 || v1 != v2 {
            return false
        }
        if !ok1 {
            return true
        }
    } 
}

func leaves(root *TreeNode, ch chan<- int)  {
    var walk func(*TreeNode)
    
    walk = func(root* TreeNode) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            ch <- root.Val
            return
        }
        walk(root.Left)
        walk(root.Right)
    }
    
    walk(root)
    close(ch)
}