func minSteps(s string, t string) int {
    var f [26]int
    var res int
    
    for _, c := range s {
        f[c - 'a']++
    }
    for _, c := range t {
        f[c - 'a']--
    }
    for _, n := range f {
        if n < 0 {
            res -= n
        } else {
            res += n
        }
    }
    return res
}