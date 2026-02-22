func maxRotateFunction(nums []int) int {
    sum := 0
    f := 0

    for i, num := range nums {
        sum += num
        f += i * num
    }

    res := f

    for i := len(nums) - 1; i > 0; i-- {
        f += sum - len(nums) * nums[i]
        res = max(res, f)
    }

    return res
}