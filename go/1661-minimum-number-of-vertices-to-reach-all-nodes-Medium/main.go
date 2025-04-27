func findSmallestSetOfVertices(n int, edges [][]int) []int {
    in := make([]bool, n)
    for _, e := range edges {
        in[e[1]] = true
    }
    var res []int
    for u := range in {
        if !in[u] {
            res = append(res, u)
        }
    }
    return res
}