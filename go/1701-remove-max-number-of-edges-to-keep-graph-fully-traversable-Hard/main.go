const (
    aliceOnly = 1
    bobOnly   = 2
    both      = 3
)

func maxNumEdgesToRemove(n int, edges [][]int) int {
    a := makeDsu(n + 1)
    b := makeDsu(n + 1)
    res := 0
    for _, e := range edges {
        u, v := e[1], e[2]
        if e[0] == both {
            if a.find(u) != a.find(v) || b.find(u) != b.find(v) {
                a.union(u, v)
                b.union(u, v)
            } else {
                res++
            }
        }
    }
    for _, e := range edges {
        u, v := e[1], e[2]
        switch e[0] {
        case aliceOnly:
            if a.find(u) != a.find(v) {
                a.union(u, v)
            } else {
                res++
            }
        case bobOnly:
            if b.find(u) != b.find(v) {
                b.union(u, v)
            } else {
                res++
            }
        }
    }
    if a.groups() != 2 || b.groups() != 2 {
        return -1
    }
    return res
}

type dsu struct {
    group []int
    size []int
}

func makeDsu(n int) dsu {
    group := make([]int, n)
    size := make([]int, n)
    for x := 0; x < n; x++ {
        group[x] = x
        size[x] = 1
    }
    return dsu{group, size}
}

func (d *dsu) union(x, y int) {
    x = d.find(x)
    y = d.find(y)
    if x == y {
        return
    }
    if d.size[x] < d.size[y] {
        d.group[x] = y
        d.size[y] += d.size[x]
    } else {
        d.group[y] = x
        d.size[x] += d.size[y]
    }
}

func (d *dsu) find(x int) int {
    for d.group[x] != x {
        x, d.group[x] = d.group[x], d.group[d.group[x]]
    }
    return x
}

func (d *dsu) groups() int {
    res := 0
    for x := range d.group {
        if x == d.group[x] {
            res++
        }
    }
    return res
}