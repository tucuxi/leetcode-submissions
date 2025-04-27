const n = 26

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    var dist [n][n]int
    for i := range n {
        for j := range n {
            dist[i][j] = math.MaxInt
        }
        dist[i][i] = 0
    }
    for i := range cost {
        a := original[i] - 'a'
        b := changed[i] - 'a'
        dist[a][b] = min(dist[a][b], cost[i])
    }
    for k := range n {
        for i := range n {
            if dist[i][k] < math.MaxInt {
               for j := range n {
                    if dist[k][j] < math.MaxInt {
                        dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
                    }
                }
            }
        }
    }

    var res int64
    for i := range source {
        a := source[i] - 'a'
        b := target[i] - 'a'
        d := dist[a][b]
        if d == math.MaxInt {
            return -1
        }
        res += int64(d)
    }
    return res
}