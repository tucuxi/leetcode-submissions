func removeAnagrams(words []string) []string {
    res := []string{words[0]}
    for i := 1; i < len(words); i++ {
        if !isAnagram(res[len(res) - 1], words[i]) {
            res = append(res, words[i]) 
        }
    }
    return res
}

func isAnagram(w1, w2 string) bool {
    if len(w1) != len(w2) {
        return false
    }
    f := [26]int{}
    for i := range w1 {
        f[w1[i] - 'a']++
        f[w2[i] - 'a']--
    }
    for i := range f {
        if f[i] != 0 {
            return false
        }
    }
    return true
}