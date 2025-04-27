func minimumFuelCost(roads [][]int, seats int) int64 {
    adj := makeGraph(roads)
    var fuel int64
    var dfs func(int, int) int64
    dfs = func(current int, previous int) int64 {
        var total int64 = 1
        for _, city := range adj[current] {
            if city != previous {
                reps := dfs(city, current)
                total += reps
                fuel += (reps + (int64(seats) - 1)) / int64(seats)
            }
        }
        return total
    }
    if len(adj) > 0 {
        dfs(0, -1)
    }
    return fuel
}

func makeGraph(roads [][]int) [][]int {
    n := 0
    for _, r := range roads {
        if r[0] >= n {
            n = r[0] + 1
        }
        if r[1] >= n {
            n = r[1] + 1
        }
    }
    adj := make([][]int, n)
    for _, r := range roads {
        adj[r[0]] = append(adj[r[0]], r[1])
        adj[r[1]] = append(adj[r[1]], r[0])
    }
    return adj
}