func findPrefixScore(nums []int) []int64 {
    conver := make([]int, 0, len(nums))
    max := 0
    for _, n := range nums {
        if n > max {
            max = n
        }
        conver = append(conver, n + max)
    }
    var sum int64
    score := make([]int64, 0, len(nums))
    for _, c := range conver {
        sum += int64(c)
        score = append(score, sum)
    }
    return score
}