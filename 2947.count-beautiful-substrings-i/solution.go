func beautifulSubstrings(s string, k int) int {
    res := 0
    for i := range s {
        v, c := 0, 0
        for j := i; j < len(s); j++ {
            if strings.IndexByte("aeiou", s[j]) == -1 {
                c++
            } else {
                v++
            }
            if v == c && v * c % k == 0 {
                res++
            }
        }
    }
    return res
}