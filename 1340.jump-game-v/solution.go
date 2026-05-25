func maxJumps(arr []int, d int) int {
    f := make([]int, len(arr))

    neighbors := func(u int) []int {
        var vs []int

        for i := u-1; i >= 0 && i >= u-d && arr[i] < arr[u]; i-- {
            vs = append(vs, i)
        }
        for i := u+1; i < len(arr) && i <= u+d && arr[i] < arr[u]; i++ {
            vs = append(vs, i)
        }
        return vs
    }

    var dfs func(int)

    dfs = func(i int) {
        if f[i] > 0 {
            return
        }
        f[i] = 1
        for _, j := range neighbors(i) {
            dfs(j)
            f[i] = max(f[i], f[j] + 1)
        }
    }

    res := 0
    
    for i := range arr {
        dfs(i)
        res = max(res, f[i])
    }

    return res
}

