func minCostConnectPoints(points [][]int) int {
    edges := make([][3]int, 0, (len(points) - 1) * (len(points) - 1) / 2)
    for i := range points {
        for j := i + 1; j < len(points); j++ {
            d := abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
            edges = append(edges, [3]int{i, j, d})
        }
    }
    sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
    u := makeUF(len(points))
    cost := 0
    for _, e := range edges {
        if u.find(e[0]) != u.find(e[1]) {
            u.union(e[0], e[1])
            cost += e[2]
        }
    }
    return cost
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
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