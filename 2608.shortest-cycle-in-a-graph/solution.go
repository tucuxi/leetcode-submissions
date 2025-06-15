func findShortestCycle(n int, edges [][]int) int {
    g := make([][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    } 
    shortest := math.MaxInt
    for v := 0; v < n; v++ {
        dist := make([]int, n)
        for q := [][]int{{v, -1, 0}}; len(q) > 0; q = q[1:] {
            w, p, l := q[0][0], q[0][1], q[0][2]
            if dist[w] > 0 {
                if dist[w] + l < shortest {
                    shortest = dist[w] + l
                } 
                break
            }
            dist[w] = l
            for _, x := range g[w] {
                if x != p {
                    q = append(q, []int{x, w, l + 1})
                }
            }
        }
    }

    if shortest == math.MaxInt {
        return -1
    }
    return shortest
}