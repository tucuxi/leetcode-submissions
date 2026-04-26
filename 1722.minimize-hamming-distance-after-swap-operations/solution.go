func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source)
    u := NewUnionFind(n)

    for _, swap := range allowedSwaps {
        u.union(swap[0], swap[1])
    }

    s := make([]map[int]int, n)

    for i := range n {
        p := u.find(i)
        if s[p] == nil {
            s[p] = make(map[int]int)
        }
        s[p][source[i]]++
    }

    res := 0

    for i := range n {
        p := u.find(i)
        if s[p][target[i]] > 0 {
            s[p][target[i]]--
        } else {
            res++
        }
    }

    return res
}

type UnionFind struct {
    parent []int
    rank   []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := range n {
        parent[i] = i
    }
    return &UnionFind{parent, rank}
}

func (u *UnionFind) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
}

func (u *UnionFind) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    if x == y {
        return
    }
    if u.rank[x] < u.rank[y] {
        x, y = y, x
    }
    u.parent[y] = x
    if u.rank[x] == u.rank[y] {
        u.rank[x]++
    }
}