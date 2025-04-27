func maximumGap(nums []int) int {
    for b := 1; b != 0; b <<= 1 {
        var p, q []int
        for _, n := range nums {
            if n & b == 0 {
                p = append(p, n)
            } else {
                q = append(q, n)
            }
        }
        i := 0
        for _, n := range p {
            nums[i] = n
            i++
        }
        for _, n := range q {
            nums[i] = n
            i++
        }
    }
    res := 0
    for i := 1; i < len(nums); i++ {
        if d := nums[i] - nums[i - 1]; d > res {
            res = d
        }
    }
    return res
}