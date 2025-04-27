func maxLengthBetweenEqualCharacters(s string) int {
    h := make(map[rune]int)
    res := -1
    for i, c := range s {
        if j, exists := h[c]; exists {
            if l := i - j - 1; l > res {
                res = l
            }
        } else { 
            h[c] = i
        }
    }
    return res
}