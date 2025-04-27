/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    levels := [][]int{}
    
    var rec func(node *Node, k int)
    rec = func(node *Node, k int) {
        if node == nil {
            return
        }
        if k == len(levels) {
            levels = append(levels, []int{ node.Val })
        } else {
            levels[k] = append(levels[k], node.Val)
        }
        for _, c := range node.Children {
            rec(c, k+1)
        }
    }
    
    rec(root, 0)
    return levels
}