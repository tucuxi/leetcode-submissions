func minExtraChar(s string, dictionary []string) int {
    dict := make(map[string]bool)
    longest := 0
    for _, w := range dictionary {
        dict[w] = true
        if len(w) > longest {
            longest = len(w)
        }
    }
    memo := make(map[int]int)
    
    var dp func(int) int
    dp = func(i int) int {
        if i == len(s) {
            return 0
        }
        if extras, exists := memo[i]; exists {
            return extras
        }
        extras := 1 + dp(i + 1)
        for j := min(len(s), i + longest); j > i; j-- {
            if _, exists := dict[s[i:j]]; exists {
                extras = min(extras, dp(j))
            }
        }
        memo[i] = extras
        return extras
    }
    
    return dp(0)
}