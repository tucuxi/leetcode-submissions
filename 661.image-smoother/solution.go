func imageSmoother(img [][]int) [][]int {
    res := make([][]int, len(img))
    for i := range img {
        res[i] = make([]int, len(img[i]))
        for j := range img[i] {
            s := 0
            c := 0
            p2 := min(len(img)-1, i+1)
            for p := max(0, i-1); p <= p2 ; p++ {
                q2 := min(len(img[i])-1, j+1)
                for q := max(0, j-1); q <= q2; q++ {
                    s += img[p][q]
                    c++
                }
            }
            res[i][j] = s/c
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}