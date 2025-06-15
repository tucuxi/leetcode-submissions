func makeConnected(n int, connections [][]int) int {
    if len(connections) < n - 1 {
        return -1
    }
    u := newUf(n)
    for _, c := range connections {
        u.union(c[0], c[1])
    }
    return u.count() - 1
}

type uf []int

func newUf(n int) uf {
    u := make(uf, n)
    for i := 0; i < n; i++ {
        u[i] = i
    }
    return u
}

func (u uf) union(a, b int) {
    a = u.find(a)
    b = u.find(b)
    u[a] = b
}

func (u uf) find(a int) int {
    for a != u[a] {
        a, u[a] = u[a], u[u[a]]
    }
    return a
}

func (u uf) count() int {
    c := 0
    for i := 0; i < len(u); i++ {
        if u.find(i) == i {
            c++
        }
    }
    return c
}