func canAliceWin(n int) bool {
    a := false
    t := 10
    for n >= t {
        n -= t
        t--
        a = !a 
    }
    return a
}