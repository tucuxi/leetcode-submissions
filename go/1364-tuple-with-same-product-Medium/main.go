func tupleSameProduct(nums []int) int {
    h := make(map[int]int)
    
    for i, a := range nums {
        for j := i + 1; j < len(nums); j++ {
            h[a * nums[j]]++
        }
    }

    res := 0

    for _, count := range h {
        pairs := (count - 1) * count / 2
        res += 8 * pairs
    }

    return res
}