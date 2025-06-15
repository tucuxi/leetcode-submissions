func executeInstructions(n int, startPos []int, s string) []int {
    res := make([]int, len(s))
    for i := range s {
        y, x := startPos[0], startPos[1]
        j := i
        loop: for ; j < len(s); j++ {
            switch s[j] {
            case 'L':
                if x == 0 {
                    break loop
                }
                x--
            case 'R':
                if x == n - 1 {
                    break loop
                }
                x++
            case 'U':
                if y == 0 {
                    break loop
                }
                y--
            case 'D':
                if y == n - 1 {
                    break loop
                }
                y++
            }
        }
        res[i] = j - i
    }
    return res
}