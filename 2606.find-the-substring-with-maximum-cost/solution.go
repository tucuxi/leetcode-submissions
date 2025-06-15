func maximumCostSubstring(s string, chars string, vals []int) int {
    v := [26]int{}
    for i := range v {
        v[i] = i + 1
    }
    for i, c := range chars {
        v[int(c - 'a')] = vals[i]
    }
    max, run := 0, 0
    for i := range s {
        if next := run + v[int(s[i] - 'a')]; next > 0 {
            run = next
            if run > max {
                max = run
            }
        } else {
            run = 0
        }
    }
    return max
}