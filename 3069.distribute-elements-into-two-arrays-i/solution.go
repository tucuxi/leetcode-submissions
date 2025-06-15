func resultArray(nums []int) []int {
    arr1 := []int{nums[0]}
    arr2 := []int{nums[1]}

    for _, n := range nums[2:] {
        if arr1[len(arr1) - 1] > arr2[len(arr2) - 1] {
            arr1 = append(arr1, n)
        } else {
            arr2 = append(arr2, n)
        }
    }

    copy(nums, arr1)
    copy(nums[len(arr1):], arr2)
    return nums
}