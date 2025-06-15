func sortEvenOdd(nums []int) []int {
    odd := make([]int, 0, len(nums) / 2)
    for i := 1; i < len(nums); i += 2 {
        odd = append(odd, nums[i])
    }
    sort.Slice(odd, func(i, j int) bool { return odd[i] > odd[j] })
    even := make([]int, 0, len(nums) / 2 + 1)
    for i := 0; i < len(nums); i += 2 {
        even = append(even, nums[i])
    }
    sort.Slice(even, func(i, j int) bool { return even[i] < even[j] })
    for i := range nums {
        if i % 2 == 0 {
            nums[i] = even[i / 2]
        } else {
            nums[i] = odd[i / 2]
        }
    }
    return nums
}