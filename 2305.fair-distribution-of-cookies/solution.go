func distributeCookies(cookies []int, k int) int {
    children := make([]int, k)

    var solve func(int) int
    solve = func(ci int) int {
        if ci == len(cookies) {
            max := 0
            for _, c := range children {
                if c > max {
                    max = c
                }
            }
            return max
        }

        c := cookies[ci]
        min := math.MaxInt
        for i := range children {
            children[i] += c
            if u := solve(ci + 1); u < min {
                min = u
            }
            children[i] -= c
        }

        return min
    }

	return solve(0)
}