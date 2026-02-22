func binaryGap(n int) int {
    res := 0
    p, q := 0, -1

    for ; n > 0; n >>=1 {
        if n & 1 != 0 {
            if q != -1 {
                res = max(res, p - q)
            }
            q = p
        }
        p++
    }

    return res
}