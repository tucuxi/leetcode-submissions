func findCircleNum(isConnected [][]int) int {
    u := newUF(len(isConnected))
    for i := 0; i < len(isConnected); i++ {
        for j := 0; j < len(isConnected[i]); j++ {
            if isConnected[i][j] == 1 {
                u.union(i, j)
            }
        }
    }
    return u.count()
}

type uf struct {
    parent []int
    size []int
}

func newUF(n int) *uf {
    parent := make([]int, n)
    size := make([]int, n)
    for i := range parent {
        parent[i] = i
        size[i] = 1
    }
    return &uf{parent, size}
}

func (u *uf) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    switch {
    case x == y:
        return
    case u.size[x] < u.size[y]:
        u.parent[x] = y
        u.size[y] += u.size[x]
    default:
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

func (u *uf) count() int {
    res := 0
    for x := range u.parent {
        if x == u.parent[x] {
            res++
        }
    }
    return res
}