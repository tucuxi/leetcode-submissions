func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    var fs, ft [26]int
    
    for i := range s {
        fs[s[i] - 'a']++
        ft[t[i] - 'a']++
    }
    return fs == ft
}