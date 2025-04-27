func maxUniqueSplit(s string) int {
    v := make(map[string]bool)

    var dfs func(int) int
    dfs = func(index int) int {
        if index == len(s) {
            return 0
        }
        c := 0
        for i := index + 1; i <= len(s); i++ {
            t := s[index:i]
            if !v[t] {
                v[t] = true
                c = max(c, 1 + dfs(i))
                delete(v, t)
            }
        }
        return c
    }

    return dfs(0)
}