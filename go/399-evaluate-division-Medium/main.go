type node struct {
    variable string
    value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    g := make(map[string][]*node)
    for i, e := range equations {
        g[e[0]] = append(g[e[0]], &node{e[1], values[i]})
        g[e[1]] = append(g[e[1]], &node{e[0], 1 / values[i]})
    }
    
    var dfs func(string, string, map[string]bool, float64) float64
    dfs = func(a, b string, v map[string]bool, acc float64) float64 {
        if a == b {
            return acc
        }
        if !v[b] {            
            v[b] = true
            for _, next := range g[b] {
                if d := dfs(a, next.variable, v, acc / next.value); d > 0 {
                    return d
                }
            }
        }
        return -1
    }
    
    res := make([]float64, len(queries))
    for i := range queries {
        if g[queries[i][0]] == nil || g[queries[i][1]] == nil {
            res[i] = -1
        } else {
            v := make(map[string]bool)
            res[i] = dfs(queries[i][0], queries[i][1], v, 1)
        }
    }
    return res
}