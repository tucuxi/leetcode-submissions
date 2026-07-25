func rearrangeString(s string, x byte, y byte) string {
    cy := 0
    
    for i := range s {
        if s[i] == y { 
            cy++
        }
    }

    res := make([]byte, len(s))

    for i := range s {
        if s[i] == x || s[i] == y {
            if cy > 0 {
                res[i] = y
                cy--
            } else {
                res[i] = x
            }
        } else {
            res[i] = s[i]
        }
    }

    return string(res)
}