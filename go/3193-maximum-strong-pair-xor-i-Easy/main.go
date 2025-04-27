func maximumStrongPairXor(nums []int) int {
    res := 0
    for i, x := range nums {
        for _, y := range nums[i:] {
            if abs(x - y) <= min(x, y) {
                res = max(res, x ^ y)
            }
        } 
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}