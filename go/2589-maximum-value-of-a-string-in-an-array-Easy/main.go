func maximumValue(strs []string) int {
    res := 0
    for _, s := range strs {
        v, err := strconv.Atoi(s)
        if err != nil {
            v = len(s)
        }
        if v > res {
            res = v
        }
    }
    return res
}