func mostProfitablePath(edges [][]int, bob int, amount []int) int {
    n := len(edges) + 1
    adj := make([][]int, n)
    adj[0] = append(adj[0], -1)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        adj[a] = append(adj[a], b)
        adj[b] = append(adj[b], a)
    }

    var reorder func(int, int)
    reorder = func(parent, node int) {
        neighbors := adj[node]
        parentIndex := slices.Index(neighbors, parent)
        if parentIndex > 0 {
            neighbors[0], neighbors[parentIndex] = neighbors[parentIndex], neighbors[0]
        }
        for _, next := range neighbors[1:] {
            reorder(node, next)
        }
    }

    reorder(-1, 0)

    q := [][]int{{0, 0}}
    res := math.MinInt
    
    for len(q) > 0 {
        for range q {
            alice, income := q[0][0], q[0][1]
            if alice == bob {
                income += amount[alice] / 2
            } else {
                income += amount[alice]
            }
            if len(adj[alice]) == 1 {
                res = max(res, income)
            } else {
                for _, next := range adj[alice][1:] {
                    q = append(q, []int{next, income})
                }
            }
            q = q[1:]
        }
        if bob > 0 {
            amount[bob] = 0
            bob = adj[bob][0]
        }
    }

    return res
}