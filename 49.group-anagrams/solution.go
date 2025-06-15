func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]byte][]string)
    for _, s := range strs {
        var c [26]byte
        for i := range s {
            c[s[i] - 'a']++
        }
        m[c] = append(m[c], s)
    }
    
    var res [][]string
    for _, anas := range m {
        res = append(res, anas)
    }
    return res 
}