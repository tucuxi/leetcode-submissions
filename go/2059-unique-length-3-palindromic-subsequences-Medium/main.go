func countPalindromicSubsequence(s string) int {
    count := 0
    var left [26]bool
    for i := range s {
        li := s[i] - 'a'
        if left[li] {
            continue
        }
        left[li] = true
        
        j := len(s) - 1
        for j > i && s[j] != s[i] {
            j--
        }
        
        var mid [26]bool
        for k := j - 1; k > i; k-- {
            ri := s[k] - 'a'
            if !mid[ri] {
                mid[ri] = true
                count++
            }
        }
    }
    return count
}