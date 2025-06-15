func frequencySort(s string) string {
    var h [80]int
    for i := range s {
        h[s[i] - '0']++
    }
    
    f := make([]int, len(h))
    for i := range f {
        f[i] = i
    }
    sort.Slice(f, func(i, j int) bool { return h[f[i]] > h[f[j]] })
    
    var b strings.Builder
    for i := range f {
        for j := h[f[i]]; j > 0; j-- {
            b.WriteByte(byte(f[i]) + '0')
        }
    }
    
    return b.String()
}