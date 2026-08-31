func missingInteger(nums []int) int {
    var v [51]bool

    for _, num := range nums {
        v[num] = true
    }

    j := nums[0]

    for i := 1; i < len(nums) && nums[i-1] + 1 == nums[i]; i++ {
        j += nums[i]
    }

    for j < len(v) && v[j] {
        j++
    }

    return j
}