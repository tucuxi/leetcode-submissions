func findRedundantConnection(edges [][]int) []int {
    u := makeUF(len(edges))
    for _, e := range edges {
        a := u.find(e[0])
        b := u.find(e[1])
        if a == b {
            return e
        }
        u.union(a, b)
    }
    return nil
}

type uf struct {
    parent []int
}

func makeUF(n int) uf {
    parent := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        parent[i] = i
    }
    return uf{parent} 
}

func (u *uf) union(a, b int) {
    a = u.find(a)
    b = u.find(b)
    if a < b {
        u.parent[b] = a
    } else if a > b {
        u.parent[a] = b
    }
}

func (u *uf) find(a int) int {
    for a != u.parent[a] {
        a, u.parent[a] = u.parent[a], u.parent[u.parent[a]]
    }
    return a
}