func validArrangement(pairs [][]int) [][]int {
    ingoing := make(map[int][]int)
    outgoing := make(map[int][]int)
    for _, pair := range pairs {
        u, v := pair[0], pair[1]
        ingoing[v] = append(ingoing[v], u)
        outgoing[u] = append(outgoing[u], v)
    }

    start := pairs[0][0]
    for u := range outgoing {
        if len(outgoing[u]) == len(ingoing[u]) + 1 {
            start = u
            break
        }
    }

    var reverseWalk []int

    var dfs func(int)
    dfs = func(u int) {
        for len(outgoing[u]) > 0 {
            v := outgoing[u][0]
            outgoing[u] = outgoing[u][1:]
            dfs(v)
        }
        reverseWalk = append(reverseWalk, u)
    }

    dfs(start)

    var res [][]int

    for i := len(reverseWalk) - 2; i >= 0; i-- {
        res = append(res, []int{reverseWalk[i + 1], reverseWalk[i]})
    }

    return res
}