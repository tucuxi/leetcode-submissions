func minSteps(s string, t string) int {
    var f [26]int
    for i := range s {
        f[s[i] - 'a']++
        f[t[i] - 'a']--
    }
    sum := 0
    for _, n := range f {
        if n > 0 {
            sum += n
        }
    }
    return sum
}
