func partitionString(s string) int {
    count := 1
    start := 0
    last := [26]int{}
    for i := range last {
        last[i] = -1
    }
    for i := range s {
        c := s[i] - 'a' 
        if last[c] >= start {
            count++
            start = i
        }
        last[c] = i
    }
    return count
}