func addToArrayForm(num []int, k int) []int {
    res := []int{}
    p := len(num)-1
    for k > 0 || p >= 0 {
        n := 0
        if p >= 0 {
            n = num[p]
            p--
        }
        d := (n+k) % 10
        k = (n+k) / 10
        res = append([]int{d}, res...)
    }
    return res
}