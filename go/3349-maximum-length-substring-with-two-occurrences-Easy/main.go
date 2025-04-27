func maximumLengthSubstring(s string) int {
    res := 0
    h := make(map[byte][]int)
    l := 0
    for r := range s {
        occ := h[s[r]]
        if len(occ) == 2 {
            l = max(l, occ[0] + 1)
            occ[0] = occ[1]
            occ[1] = r
        } else {
            occ = append(occ, r)
            h[s[r]] = occ
        }
        res = max(res, r-l+1)
    }
    return res
}