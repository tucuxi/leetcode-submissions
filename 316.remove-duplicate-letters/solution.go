func removeDuplicateLetters(s string) string {
    last := make(map[rune]int)
    for i, c := range s {
        last[c] = i
    }
    res := make([]rune, 0, 26)
    used := make(map[rune]bool)
    for i, c := range s {
        if !used[c] {
            for r := len(res) - 1; r >= 0 && c < res[r] && last[res[r]] > i; r-- {
                used[res[r]] = false
                res = res[:r]
            }
            used[c] = true
            res = append(res, c)
        }
    }
    return string(res)
}