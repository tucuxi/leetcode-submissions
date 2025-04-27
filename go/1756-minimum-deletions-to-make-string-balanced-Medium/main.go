func minimumDeletions(s string) int {
    a := 0
    b := 0
    res := len(s)

    for _, ch := range s {
        if ch == 'a' {
            a++
        }
    }
    for _, ch := range s {
        if ch == 'a' {
            a--
        } 
        res = min(res, a + b)
        if ch == 'b' {
            b++
        }
    }
    return res
}