func lexicalOrder(n int) []int {
    var res []int
    var dfs func(k int) 

    dfs = func(k int) {
        if k <= n {
            res = append(res, k)
            for i := range 10 {
                dfs(10 * k + i)
            }
        }
    }

    for i := 1; i < 10; i++ {
        dfs(i)
    }
    return res
}