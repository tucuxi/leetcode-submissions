type entry struct {
    node int
    prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    outgoing := make([][]entry, n)
    for i, e := range edges {
        outgoing[e[0]] = append(outgoing[e[0]], entry{e[1], succProb[i]})
        outgoing[e[1]] = append(outgoing[e[1]], entry{e[0], succProb[i]})
    }
    prob := make([]float64, n)
    prob[start] = 1.0
    queue := []int{start}
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        p := prob[u]
        for _, e := range outgoing[u] {
            v := e.node
            q := e.prob
            if prob[v] < p * q {
                prob[v] = p * q
                queue = append(queue, v)
            }
        }
    }
    return prob[end]
}