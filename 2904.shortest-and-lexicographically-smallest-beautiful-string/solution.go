func shortestBeautifulSubstring(s string, k int) string {
    if strings.Count(s, "1") < k {
        return ""
    }
    
    res := s
    ones := 0
    left := 0

    for right := range s {
        if s[right] == '1' {
            ones++
        }
        for ; ones > k || s[left] == '0'; left++ {
            if s[left] == '1' {
                ones--
            }
        }
        if ones == k {
            t := s[left : right+1]
            if len(t) < len(res) || len(t) == len(res) && t < res {
                res = t
            }
        }
    }
    return res
}