func kthLargestNumber(nums []string, k int) string {
    return quickselect(nums, 0, len(nums) - 1, k - 1)
}

func partition(nums []string, left, right, pivot int) int {
    pivotValue := nums[pivot]
    nums[pivot], nums[right] = nums[right], pivotValue
    store := left
    for i := left; i < right; i++ {
        if greater(nums[i], pivotValue) {
            nums[store], nums[i] = nums[i], nums[store]
            store++
        }
    }
    nums[store], nums[right] = nums[right], nums[store]
    return store
}

func greater(num1, num2 string) bool {
    if len(num1) != len(num2) {
        return len(num1) > len(num2)
    }
    for i := range num1 {
        if num1[i] != num2[i] {
            return num1[i] > num2[i]
        }
    }
    return false
}

func quickselect(nums []string, left, right, k int) string {
    for left < right {
        pivot := partition(nums, left, right, left + rand.Intn(right - left))
        if k == pivot {
            return nums[k]
        } else if k < pivot {
            right = pivot - 1
        } else {
            left = pivot + 1
        }
    }
    return nums[right]
}