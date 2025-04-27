type uf []int

func newUf(n int) uf {
    u := make(uf, n)
    for i := range u {
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

func minScore(n int, roads [][]int) int {
    u := newUf(n)
    for _, r := range roads {
        u.union(r[0] - 1, r[1] - 1)
    }
    p := u.find(0)
    res := math.MaxInt
    for _, r := range roads {
        if u.find(r[0] - 1) == p || u.find(r[1] - 1) == p {
            if r[2] < res {
                res = r[2]
            }
        }
    }
    return res
}