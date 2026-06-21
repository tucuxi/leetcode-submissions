func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
    g := make([][]bool, n)
    
    for i := range n {
        g[i] = make([]bool, n)
    }

    for _, d := range dig {
        r, c := d[0], d[1]
        g[r][c] = true
    }

    res := 0

    outer:
    for _, a := range artifacts {
        r1, c1, r2, c2 := a[0], a[1], a[2], a[3]
        for r := r1; r <= r2; r++ {
            for c := c1; c <= c2; c++ {
                if !g[r][c] {
                    continue outer
                }
            }
        }
        res++
    }

    return res
}