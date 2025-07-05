func maxDistance(s string, k int) int {
    res := 0
    h, v := 0, 0
    for i := range s {
        switch s[i] {
            case 'N': v--
            case 'S': v++
            case 'W': h--
            case 'E': h++
        }
        res = max(res, min(abs(h) + abs(v) + 2 * k, i + 1))
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}