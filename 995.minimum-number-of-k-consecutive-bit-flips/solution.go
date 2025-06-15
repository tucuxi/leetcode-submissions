func minKBitFlips(nums []int, k int) int {
    var (
        flipped bool
        flips []bool
        res int
    )

    for i := range nums {
        if i >= k {
            flipped = flipped != flips[0]
            flips = flips[1:]
        }
        if flipped == (nums[i] == 1) {
            if i + k > len(nums) {
                return -1
            }
            flipped = !flipped
            flips = append(flips, true)
            res++
        } else {
            flips = append(flips, false)
        }
    }
    return res
}