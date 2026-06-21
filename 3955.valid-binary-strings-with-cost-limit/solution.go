func generateValidStrings(n int, k int) []string {
    var (
        res []string
        dfs func(string, int)
    )

    dfs = func(pre string, cost int) {
        i := len(pre)
        if i == n {
            res = append(res, pre)
            return
        }
        dfs(pre + "0", cost)
        if c := cost+i; c <= k && (i == 0 || pre[i-1] == '0') {
            dfs(pre + "1", c)
        }        
    }

    dfs("", 0)
    return res
}