func findDuplicate(nums []int) int {
    dup := -1
    lo := 1
    hi := len(nums) - 1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        count := 0
        for _, n := range nums {
            if n <= mid {
                count++
            }
        }
        if count <= mid {
            lo = mid + 1
        } else {
            dup = mid
            hi = mid - 1
        }
    }
    return dup
}