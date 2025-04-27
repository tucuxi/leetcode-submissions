func maxJump(stones []int) int {
    max := stones[1] - stones[0]
    for i := 2; i < len(stones); i++ {
        if l := stones[i] - stones[i - 2]; l > max {
            max = l
        }
    }
    return max
}