func minimumSteps(s string) int64 {
    var res, ones int64
    for _, r := range s {
        if r == '1' {
            ones++
        } else {
            res += ones
        }
    }
    return res
}