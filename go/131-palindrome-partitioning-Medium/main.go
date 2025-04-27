func partition(s string) [][]string {
    var res [][]string

    var dfs func(string, []string)
    dfs = func(s string, parts []string) {
        if len(s) == 0 {
            solution := make([]string, len(parts))
            copy(solution, parts)
            res = append(res, solution)
            return
        }
        for i := 1; i <= len(s); i++ {
            if isPalindrome(s[:i]) {
                dfs(s[i:], append(parts, s[:i]))
            }
        }
    }

    dfs(s, make([]string, 0, len(s)))
    return res
}

func isPalindrome(s string) bool {
    for p, q := 0, len(s) - 1; p < q; p, q = p + 1, q - 1 {
        if s[p] != s[q] {
            return false
        }
    }
    return true
}