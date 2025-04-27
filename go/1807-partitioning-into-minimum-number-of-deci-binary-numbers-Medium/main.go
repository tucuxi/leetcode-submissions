func minPartitions(n string) int {
    res := 0
    for _, d := range n {
        m := int(d - '0')
        if m > res {
            res = m
        }
    }
    return res
}