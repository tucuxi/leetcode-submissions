func finalString(s string) string {
    b := make([]byte, 0, len(s))
    for i := range s {
        if s[i] == 'i' {
            for j, k := 0, len(b) - 1; j < k; j, k = j + 1, k - 1 {
                b[j], b[k] = b[k], b[j]
            } 
        } else {
            b = append(b, s[i])
        }
    }
    return string(b)
}