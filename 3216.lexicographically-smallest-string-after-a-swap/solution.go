func getSmallestString(s string) string {
    b := []byte(s)

    for i := 1; i < len(b); i++ {
        if b[i-1] & 1 == b[i] & 1 && b[i-1] > b[i] {
            b[i-1], b[i] = b[i], b[i-1]
            break
        } 
    }

    return string(b)
}