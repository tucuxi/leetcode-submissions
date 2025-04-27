/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    h := map[*Node]*Node{}

    var dfs func(*Node) *Node
    dfs = func(node *Node) *Node {
        if clone := h[node]; clone != nil {
            return clone
        }
        clone := &Node{Val: node.Val}
        h[node] = clone
        for _, neighbor := range node.Neighbors {
            neighborClone := dfs(neighbor)
            clone.Neighbors = append(clone.Neighbors, neighborClone)
        }
        return clone
    }

    if node == nil {
        return nil
    }
    return dfs(node)
}