func reverseWords(s string) string {
    var b strings.Builder
    for i := 0; i < len(s); {
        if s[i] == ' ' {
            b.WriteByte(' ')
            i++
        } else {
            j := i + 1
            for j < len(s) && s[j] != ' ' {
                j++
            }
            for k := j - 1; k >= i; k-- {
                b.WriteByte(s[k])
            }
            i = j
        }
        
    }
    return b.String()
}