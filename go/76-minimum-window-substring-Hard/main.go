func minWindow(s string, t string) string {
    ft := make(map[byte]int)
    cnt := 0
    for i := range t {
        ft[t[i]]++
        cnt++
    }
    
    var minSeq string
    minLength := math.MaxInt
    p, q := 0, 0
    
    for q < len(s) {
        if ft[s[q]] > 0 {
            cnt--
        }
        ft[s[q]]--
        q++
        for ; cnt == 0; p++ {
            if q - p < minLength {
                minLength = q - p
                minSeq = s[p: p + minLength]
            }
            ft[s[p]]++
            if ft[s[p]] > 0 {
                cnt++
            }
        }
    }
    return minSeq
}