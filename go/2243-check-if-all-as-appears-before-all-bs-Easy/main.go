func checkString(s string) bool {
    b := false
    for _, c := range s {
        if (c == 'b') {
            b = true
        } else if b {
            return false
        }
        
    }
    return true
}