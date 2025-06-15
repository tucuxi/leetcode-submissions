func sortVowels(s string) string {
    var vowels []byte
    
    for i := range s {
        if isVowel(s[i]) {
            vowels = append(vowels, s[i])
        }
    }
    sort.Slice(vowels, func(i, j int) bool { return vowels[i] < vowels[j] })
    
    res := []byte(s)
    k := 0 
    for i := range s {
        if isVowel(s[i]) {
            res[i] = vowels[k]
            k++
        }
    }
    return string(res)
}

func isVowel(c byte) bool {
    return strings.IndexByte("aeiouAEIOU", c) != -1
}