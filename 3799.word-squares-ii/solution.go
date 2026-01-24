func wordSquares(words []string) [][]string {
    var res [][]string
    
    for _, top := range words {
        for _, left := range words {
            if left == top || top[0] != left[0] {
                continue
            }
            for _, bottom := range words {
                if bottom == top || bottom == left || bottom[0] != left[3] {
                    continue
                }
                for _, right := range words {
                    if right == top || right == left || right == bottom || right[0] != top[3] || right[3] != bottom[3] {
                        continue
                    }
                    res = append(res, []string{top, left, right, bottom})
                }
            }
        }
    }
    slices.SortFunc(res, func(a, b []string) int {
        for i := range a {
            if a[i] < b[i] {
                return -1
            }
            if a[i] > b[i] {
                return 1
            }
        }
        return 0
    })
    return res
}