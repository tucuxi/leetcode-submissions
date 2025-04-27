func isArraySpecial(nums []int, queries [][]int) []bool {
    f := []int{0}
    k := 0

    for i := 1; i < len(nums); i++ {
        if (nums[i-1] & 1) == (nums[i] & 1) {
            k++
        }
        f = append(f, k)
    }
    
    var res []bool

    for _, q := range queries {
        res = append(res, f[q[0]] == f[q[1]])     
    }
    
    return res
}