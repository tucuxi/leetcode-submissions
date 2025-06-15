func numberOfSpecialChars(word string) int {
    h := make(map[rune]bool)
    res := 0
    for _, r := range word {
        if !h[r] {
            h[r] = true
            if unicode.IsLower(r){
                if h[unicode.ToUpper(r)] {
                    res++
                }
            } else {
                if h[unicode.ToLower(r)] {
                    res++
                }                
            }
        }
    }
    return res
}