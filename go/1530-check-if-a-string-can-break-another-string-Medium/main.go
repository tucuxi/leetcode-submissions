func checkIfCanBreak(s1 string, s2 string) bool {
    x, y := []byte(s1), []byte(s2)
    slices.Sort(x)
    slices.Sort(y)
    lt, gt := false, false
    for i := range x {
        if x[i] < y[i] {
            if gt {
                return false
            }
            lt = true
        } else if x[i] > y[i] {
            if lt {
                return false
            }
            gt = true
        }
    }
    return true
}