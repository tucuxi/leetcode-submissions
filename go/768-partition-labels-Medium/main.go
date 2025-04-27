func partitionLabels(s string) []int {
    last := make(map[byte]int)
    for i := range s {
        last[s[i]] = i
    }
    
    var res []int
    start, end := 0, 0
    for i := range s {
        if pos, ok := last[s[i]]; ok {
            if pos > end {
                end = pos
            }
        }
        if end == i {
            res = append(res, end - start + 1)
            start = end + 1
        }
    }
    return res
}