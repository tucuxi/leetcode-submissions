func findTheDifference(s string, t string) byte {
    var h [26]int
    for _, a := range s {
        h[int(a) - 'a']--
    }
    for _, a := range t {
        h[int(a) - 'a']++
    }
    for a, f := range h {
        if f > 0 {
            return byte(a) + 'a'
        }
    }
    panic("no solution")
}