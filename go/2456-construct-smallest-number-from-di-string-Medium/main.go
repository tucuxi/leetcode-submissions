func smallestNumber(pattern string) string {
    v := make(map[rune]bool)
    b := make([]rune, 0, len(pattern) + 1)
    res := make([]rune, len(pattern) + 1)
    for i := range res {
        res[i] = '9'
    }
    
    var rec func(int)
    rec = func(index int) {
        if index == len(pattern) {
            if slices.Compare(b, res) < 0 {
                copy(res, b)
            } 
            return
        }
        for ch := '1'; ch <= '9'; ch++ {
            if v[ch] {
                continue
            }
            if pattern[index] == 'I' && ch < b[index] {
                continue
            }
            if pattern[index] == 'D' && ch > b[index] {
                continue
            }
            v[ch] = true
            b = append(b, ch)
            rec(index + 1)
            delete(v, ch)
            b = b[:len(b) - 1]
        }
    }
    
    for ch := '1'; ch <= '9'; ch++ {
        clear(v)
        b = b[:0]
        v[ch] = true
        b = append(b, ch)
        rec(0)
    }
    return string(res)
}