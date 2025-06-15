func averageValue(nums []int) int {
    sum, cnt := 0, 0
    for _, n := range nums {
        if n % 6 == 0 {
            sum += n
            cnt++
        }
    }
    if cnt == 0 {
        return 0
    }
    return sum / cnt
}