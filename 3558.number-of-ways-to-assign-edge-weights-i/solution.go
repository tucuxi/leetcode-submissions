const mod = 1000000007

func assignEdgeWeights(edges [][]int) int {
    n := len(edges) + 1
    adj := make([][]int, n+1)
 
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    
    var dfs func(int, int) int

    dfs = func(u, p int) int {
        d := 0
        for _, v := range adj[u] {
            if v != p {
                d = max(d, dfs(v, u) + 1)
            }
        }
        return d
    }

    d := dfs(1, 0)
    return pow2(d-1)
}

func pow2(n int) int {
	res := 1
    b := 2
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * b % mod
		}
		b = b * b % mod
	}
	return res
}