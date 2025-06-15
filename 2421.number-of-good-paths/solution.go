func numberOfGoodPaths(vals []int, edges [][]int) int {
    n := len(vals)
    root := make([]int, n)
    count := make([]map[int]int, n)
    for i := 0; i < n; i++ {
        root[i] = i
        count[i] = map[int]int{ vals[i]: 1 }
    }
    sort.Slice(edges, func(i, j int) bool {
        return max(vals[edges[i][0]], vals[edges[i][1]]) < max(vals[edges[j][0]], vals[edges[j][1]])
    })
    res := n
    for _, e := range edges {
        i := find(root, e[0])
        j := find(root, e[1])
        v := max(vals[e[0]], vals[e[1]])
        ci := count[i][v]
        cj := count[j][v]
        res += ci * cj
        root[j] = i
        count[i][v] = ci + cj

    }
    return res
}

func find(root []int, x int) int {
    for x != root[x] {
        x, root[x] = root[x], root[root[x]]
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}