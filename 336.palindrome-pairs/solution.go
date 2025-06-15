func palindromePairs(words []string) [][]int {
    d := map[string]int{}
    for i, s := range words {
        d[s] = i
    }
    res := [][]int{}
    for i, s := range words {
        if k, ok := d[rev(s)]; ok {
            if i != k {
                res = append(res, []int{i, k})
            }
        }
        if len(s) > 0 && pal(s) {
            if k, ok := d[""]; ok {
                res = append(res, []int{i, k}, []int{k, i})
            }
        }
        for j := 1; j < len(s); j++ {
            if pal(s[:j]) {
                if k, ok := d[rev(s[j:])]; ok {
                    res = append(res, []int{k, i})
                }
            }
            if pal(s[j:]) {
                if k, ok := d[rev(s[:j])]; ok {
                    res = append(res, []int{i, k})
                }
            }
        }
    }
    return res
}

func pal(s string) bool {
    for i, j := 0, len(s) - 1; i < j; {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func rev(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}