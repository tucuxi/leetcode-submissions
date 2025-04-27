func minFlips(a int, b int, c int) int {
    res := 0
    for bit := 1; bit != 0; bit <<= 1 {
        if c & bit == 0 {
            if a & bit != 0 {
                res++
            }
            if b & bit != 0 {
                res++
            }
        } else {
            if (a | b) & bit == 0 {
                res++
            }
        }
    }
    return res
}