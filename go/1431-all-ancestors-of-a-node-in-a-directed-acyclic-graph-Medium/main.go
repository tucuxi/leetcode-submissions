func getAncestors(n int, edges [][]int) [][]int {
    indegree := make([]int, n)
    outgoing := make([][]int, n)

    for _, e := range edges {
        from, to := e[0], e[1]
        indegree[to]++
        outgoing[from] = append(outgoing[from], to)
    }

    var zeroIndegreeNodes []int
    for i := range n {
        if indegree[i] == 0 {
            zeroIndegreeNodes = append(zeroIndegreeNodes, i)
        }
    }

    ancestors := make([]map[int]bool, n)
    for len(zeroIndegreeNodes) > 0 {
        var newZeroIndegreeNodes []int
        for _, i := range zeroIndegreeNodes {
            for _, j := range outgoing[i] {
                if ancestors[j] == nil {
                    ancestors[j] = map[int]bool{i: true}
                } else {
                    ancestors[j][i] = true
                }
                for k := range ancestors[i] {
                    ancestors[j][k] = true
                }
                indegree[j]--
                if indegree[j] == 0 {
                    newZeroIndegreeNodes = append(newZeroIndegreeNodes, j)
                }
            }
        }
        zeroIndegreeNodes = newZeroIndegreeNodes
    }

    res := make([][]int, n)

    for i := range res {
        for j := range ancestors[i] {
            res[i] = append(res[i], j)
        }
        slices.Sort(res[i])
    }

    return res
}