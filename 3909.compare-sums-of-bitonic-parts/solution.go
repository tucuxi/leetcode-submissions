func compareBitonicSums(nums []int) int {
    n := len(nums)
    sumA := nums[0]
    i := 1

    for i < n && nums[i] > nums[i-1] {
        sumA += nums[i]
        i++
    }

    sumD := nums[i-1]
    for i < n {
        sumD += nums[i]
        i++
    }

    if sumA > sumD {
        return 0
    }
    if sumA < sumD {
        return 1
    }
    return -1
}