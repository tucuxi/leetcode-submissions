func minCost(nums []int, queries [][]int) []int {
    n := len(nums)
    f := make([]int, n)
    b := make([]int, n)

    for i := 0; i < n-1; i++ {
        if i == 0 || nums[i] - nums[i-1] > nums[i+1] - nums[i] {
            f[i+1] = f[i] + 1 
        } else {
            f[i+1] = f[i] + nums[i+1] - nums[i]
        }
    }
    for i := 1; i < n; i++ {
        if i == n-1 || nums[i] - nums[i-1] <= nums[i+1] - nums[i] {
            b[i] = b[i-1] + 1
        } else {
            b[i] = b[i-1] + nums[i] - nums[i-1]
        }
    }
    
    res := make([]int, 0, len(queries))

    for _, q := range queries {
        l, r := q[0], q[1]
        if l < r {
            res = append(res, f[r] - f[l])
        } else {
            res = append(res, b[l] - b[r])
        }
    }

    return res
}