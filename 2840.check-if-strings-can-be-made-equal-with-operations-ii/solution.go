func checkStrings(s1 string, s2 string) bool {
    return check(s1, s2, 0) && check(s1, s2, 1)
}

func check(s1 string, s2 string, start int) bool {
    var h [26]int

    for i := start; i < len(s1); i += 2 {
        h[s1[i] - 'a']++
        h[s2[i] - 'a']--
    }
    for _, c := range h {
        if c != 0 {
            return false
        }
    }
    return true
}