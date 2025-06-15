func appealSum(s string) int64 {
    var (
        res, sum int64 
        occ [26]int
    )
    for i := range occ {
        occ[i] = -1
    }
    for i := range s {
        k := s[i] - 'a'
        sum += int64(i - occ[k])
        res += sum
        occ[k] = i
    }
    return res
}