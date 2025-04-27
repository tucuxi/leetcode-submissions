func minimumCost(n int, edges [][]int, query [][]int) []int {
    item := make([]int, n)
    rank := make([]int, n)
    cost := make([]int, n)
    for i := range item {
        item[i] = i
        rank[i] = 1
        cost[i] = -1
    }
    for _, p := range edges {
        u, v := find(item, p[0]), find(item, p[1])
        if u != v {
            union(item, rank, p[0], p[1])
            switch {
            case cost[u] == -1 && cost[v] == -1:
                cost[find(item, p[0])] = p[2]
            case cost[u] == -1:
                cost[find(item, p[0])] = cost[v] & p[2]
            case cost[v] == -1:
                cost[find(item, p[0])] = cost[u] & p[2]
            default:
                cost[find(item, p[0])] = cost[u] & cost[v] & p[2]
            }
        } else {
            cost[u] &= p[2]
        }
    }
    res := make([]int, len(query))
    for i, q := range query {
        if q[0] == q[1] {
            res[i] = 0
            continue
        }
        u, v := find(item, q[0]), find(item, q[1])
        if u == v {
            res[i] = cost[u]
        } else {
            res[i] = -1
        }
    }
    return res
}

func find(item []int, i int) int {
    p, c := item[i], i
    if p != c {
        p = find(item, p)
        item[c] = p
    }
    return p
}

func union(item, rank []int, i, j int) {
    u, v := find(item, i), find(item, j)
    switch {
        case u == v:
        case rank[u] < rank[v]:
        item[u] = v
        case rank[u] > rank[v]:
        item[v] = u
        default:
        item[u] = v
        rank[v]++
    }
}