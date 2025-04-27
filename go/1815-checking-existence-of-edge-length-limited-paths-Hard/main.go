func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
    index := make([]int, len(queries))
    for i := range index {
        index[i] = i
    }
    sort.Slice(index, func(i, j int) bool { return queries[index[i]][2] < queries[index[j]][2] })
    sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })
    u := newUF(n)
    res := make([]bool, len(queries))
    k := 0
    for _, i := range index {
        limit := queries[i][2]
        for ; k < len(edgeList) && edgeList[k][2] < limit; k++  {
            u.union(edgeList[k][0], edgeList[k][1])
        }
        res[i] = u.connected(queries[i][0], queries[i][1])
    }
    return res
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

func (u *uf) connected(x, y int) bool {
    return u.find(x) == u.find(y)
}