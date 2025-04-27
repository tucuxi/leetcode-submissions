func numSimilarGroups(strs []string) int {
    u := newUF(len(strs))
    for i := range strs {
        for j := i + 1; j < len(strs); j++ {
            if u.find(i) != u.find(j) && similar(strs[i], strs[j]) {
                u.union(i, j)
            }
        }
    }
    return u.groups()
}

func similar(a, b string) bool {
    d := 0
    for i := range a {
        if a[i] != b[i] {
            d++
        }
    }
    return d == 0 || d == 2
}

type uf struct {
    parent []int
}

func newUF(n int) *uf {
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    return &uf{p}
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

func (u *uf) groups() int {
    res := 0
    for i := range u.parent {
        if i == u.parent[i] {
            res++
        }
    }
    return res
}