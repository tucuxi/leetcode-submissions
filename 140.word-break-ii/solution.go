func wordBreak(s string, wordDict []string) []string {
    h := make(map[string]bool)
    for _, w := range wordDict {
        h[w] = true
    }

    var res []string

    var dfs func(int, int, []string)

    dfs = func(left int, right int, acc []string) {
        if left == len(s) {
            var b strings.Builder
            for i := 0; i < len(acc) - 1; i++ {
                b.WriteString(acc[i])
                b.WriteString(" ")
            }
            b.WriteString(acc[len(acc) - 1])
            res = append(res, b.String())
            return
        }
        if right > len(s) {
            return
        }
        if h[s[left:right]] {
            dfs(right, right+1, append(acc, s[left:right]))
        }
        dfs(left, right+1, acc)
    }

    dfs(0, 1, nil)
    return res
}