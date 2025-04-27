func sumSubarrayMins(arr []int) int {
    const MOD = 1_000_000_007
    
    left := func() []int {
        var (
            l = make([]int, len(arr))
            s []int
        )
        for i := range l {
            l[i] = -1
        } 
        for i := range arr {
            for len(s) > 0 && arr[s[len(s) - 1]] >= arr[i] {
                s = s[:len(s) - 1]
            }
            if len(s) > 0 {
                l[i] = s[len(s) - 1]
            }
            s = append(s, i)
        }
        return l
    }()
        
    right := func() []int {
        var (
            r = make([]int, len(arr))
            s []int
        )
        for i := range r {
            r[i] = len(arr)
        }
        for i := len(arr) - 1; i >= 0; i-- {
            for len(s) > 0 && arr[s[len(s) - 1]] > arr[i] {
                s = s[:len(s) - 1]
            }
            if len(s) > 0 {
                r[i] = s[len(s) - 1]
            }
            s = append(s, i)
        }
        return r
    }()
    
    return func() int {
        res := 0
        for i := range arr {
            res = (res + (i - left[i]) * (right[i] - i) * arr[i]) % MOD
        }
        return res
    }()
}