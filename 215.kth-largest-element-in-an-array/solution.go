func findKthLargest(nums []int, k int) int {
    p := 0
    q := len(nums) - 1
    for p < q {
        i := partition(nums, p, q, p)
        if i + 1 < k {
            p = i + 1
        } else {
            q = i
        }
    }
    return nums[p]
}

func partition(nums []int, p, q, k int) int {
    pivot := nums[k]
    nums[k], nums[q] = nums[q], nums[k]
    for i := p; i < q; i++ {
        if nums[i] > pivot {
            nums[i], nums[p] = nums[p], nums[i]
            p++
        }
    }
    nums[p], nums[q] = nums[q], nums[p]
    return p
}