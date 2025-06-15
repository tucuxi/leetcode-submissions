func minMutation(startGene string, endGene string, bank []string) int {
    m := 0
    v := make([]bool, len(bank))
    q := []string{startGene}
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            gene := q[0]
            q = q[1:]
            if gene == endGene {
                return m
            }
            for j := range bank {
                if !v[j] && diff(bank[j], gene) == 1 {
                    v[j] = true
                    q = append(q, bank[j])
                }
            }
        }
        m++
    }
    return -1
}

func diff(a, b string) int {
    res := 0
    for i := range a {
        if a[i] != b[i] {
            res++
        }
    }
    return res
}