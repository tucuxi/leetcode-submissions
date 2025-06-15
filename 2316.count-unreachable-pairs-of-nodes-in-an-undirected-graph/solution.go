type uf struct {
    parent []int
    size []int
}

func makeUf(n int) uf {
    p := make([]int, n)
    s := make([]int, n)
    for i := 0; i < n; i++ {
        p[i] = i
        s[i] = 1
    }
    return uf{parent: p, size: s}
}

func (u uf) union(a, b int) {
    a = u.find(a)
    b = u.find(b)
    if a == b {
        return
    }
    if u.size[a] < u.size[b] {
        a, b = b, a
    }
    u.parent[b] = a
    u.size[a] += u.size[b]
}

func (u uf) find(a int) int {
    for a != u.parent[a] {
        a, u.parent[a] = u.parent[a], u.parent[u.parent[a]]
    }
    return a
}

func (u uf) sets() map[int] int {
    res := make(map[int]int)
    for _, i := range u.parent {
        if i == u.parent[i] {
            res[i] = u.size[i]
        }
    }
    return res
}

func countPairs(n int, edges [][]int) int64 {
    u := makeUf(n)
    for _, e := range edges {
        u.union(e[0], e[1])
    }
    var count int64
    for _, s := range u.sets() {
        count += int64(s * (n - s))
    }
    return count / 2
}

