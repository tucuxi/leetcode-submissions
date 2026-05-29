func minimumSwaps(nums []int) int {
    i := 0
    j := len(nums) - 1
    z := 0

    for i < j {
        if nums[i] == 0 {
            if nums[j] == 0 {
                j--
                continue
            }
            j--
            z++
        }
        i++
    }

    return z
}