func isBipartite(graph [][]int) bool {
    u := newUF(len(graph))
    for x, edges := range graph {
        for _, y := range edges {
            if u.find(x) == u.find(y) {
                return false
            }
            u.union(edges[0], y)
        }
    }
    return true
}

type uf struct {
    parent []int
    size []int
}

func newUF(n int) *uf {
    parent := make([]int, n)
    size := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        size[i] = 1
    }
    return &uf{parent, size}
}

func (u *uf) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    if x == y {
        return
    }
    if u.size[x] < u.size[y] {
        u.parent[x] = y
        u.size[y] += u.size[x]
    } else {
        u.parent[y] = x
        u.size[x] += u.size[y]
    }
}

func (u *uf) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
}