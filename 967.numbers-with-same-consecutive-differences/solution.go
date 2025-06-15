func numsSameConsecDiff(n int, k int) []int {
    res := []int{}
 
    var rec func(num int, rem int, prev int)
    rec = func(num int, rem int, prev int) {
        if rem == 0 {
            res = append(res, num)
            return
        }
        if next := prev - k; next >= 0 {
            rec(num * 10 + next, rem - 1, next)
        }
        if next := prev + k; next < 10 && k != 0 {
            rec(num * 10 + next, rem - 1, next)
        }
    }
    
    for i := 1; i <= 9; i++ {
        rec(i, n - 1, i)
    }
    return res
}