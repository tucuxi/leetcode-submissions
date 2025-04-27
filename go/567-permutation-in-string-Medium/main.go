func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    var f, z [26]int
    for i := range s1 {
        f[s1[i] - 'a']--
        f[s2[i] - 'a']++
    }
    if f == z {
        return true
    }
    for i := len(s1); i < len(s2); i++ {
        f[s2[i] - 'a']++
        f[s2[i - len(s1)] - 'a']-- 
        if f == z {
            return true
        }        
    }
    return false
}