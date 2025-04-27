func wiggleSort(nums []int)  {
    n := len(nums)
    median := quickselect(nums, n / 2)
    for i, j, k := 0, 0, n - 1; j <= k; {
        jp := (2 * j + 1) % (n | 1)
        if nums[jp] > median {
            ip := (2 * i + 1) % (n | 1)
            nums[jp], nums[ip] = nums[ip], nums[jp]
            i++
            j++
        } else if nums[jp] < median {
            kp := (2 * k + 1) % (n | 1)
            nums[jp], nums[kp] = nums[kp], nums[jp]
            k--
        } else {
            j++
        }
    }
}

func quickselect(nums []int, k int) int {
    i, j := 0, len(nums) - 1
    for i < j {
        pivotIndex := partition(nums, i, j, i + rand.Intn(j - i))
        if pivotIndex == k {
            return nums[k]
        } else if pivotIndex < k {
            i = pivotIndex + 1
        } else {
            j = pivotIndex - 1
        }
    }
    return nums[j]
}

func partition(nums []int, left, right, pivotIndex int) int {
    pivotValue := nums[pivotIndex]
    nums[pivotIndex] = nums[right]
    nums[right] = pivotValue
    storeIndex := left
    for i := left; i < right; i++ {
        if nums[i] < pivotValue {
            nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
            storeIndex++
        }
    }
    nums[right] = nums[storeIndex]
    nums[storeIndex] = pivotValue
    return storeIndex
}