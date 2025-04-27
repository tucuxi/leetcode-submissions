/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    res := []int{}
    adj := make(map[*TreeNode][]*TreeNode)
    v := make(map[*TreeNode]bool)
    
    var makeAdj func(*TreeNode)
    makeAdj = func(node *TreeNode) {
        if node.Left != nil {
            adj[node] = append(adj[node], node.Left)
            adj[node.Left] = append(adj[node.Left], node)
            makeAdj(node.Left)
        }
        if node.Right != nil {
            adj[node] = append(adj[node], node.Right)
            adj[node.Right] = append(adj[node.Right], node)
            makeAdj(node.Right)
        }
    }
    
    if root != nil {
        makeAdj(root)
    }
    q := []*TreeNode{target}
    v[target] = true
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            if k == 0 {
                res = append(res, q[0].Val)
            } else {
                for _, p := range adj[q[0]] {
                    if !v[p] {
                        v[p] = true
                        q = append(q, p)
                    }
                }
            }
            q = q[1:]
        }
        k--
    }
    return res
}