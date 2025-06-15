/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    var res []int
    var dfs func(*Node)

    dfs = func(node *Node) {
        if node != nil {
            for _, child := range node.Children {
                dfs(child)
            }
            res = append(res, node.Val)
        }
    }

    dfs(root)
    return res
}