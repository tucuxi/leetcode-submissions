func findMaxConsecutiveOnes(nums []int) int {
    max, run := 0, 0
    for _, num := range nums {
        if num == 1 {
            run++
            if run > max {
                max = run
            }
        } else {
            run = 0
        }
    }
    return max
}