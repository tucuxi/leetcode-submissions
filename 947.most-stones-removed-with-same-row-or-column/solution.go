func removeStones(stones [][]int) int {
    u := makeUF(len(stones))
    for i, si := range stones {
        for j, sj := range stones {
            if si[0] == sj[0] || si[1] == sj[1] {
                u.union(i, j)
            }            
        }
    }
    return len(stones) - u.groups()
}

type uf struct {
    parent []int
    size []int
}

func makeUF(n int) uf {
    parent := make([]int, n)
    size := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        size[i] = 1
    }
    return uf{parent, size}
}

func (u *uf) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    if x != y {
        if u.size[x] < u.size[y] {
            u.parent[x] = y
            u.size[y] += u.size[x]
        } else {
            u.parent[y] = x
            u.size[x] += u.size[y]
        }
    }
}

func (u *uf) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
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