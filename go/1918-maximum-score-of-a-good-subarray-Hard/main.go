func maximumScore(nums []int, k int) int {
    left := k
    right := k
    currMin := nums[k]
    res := nums[k]
    for left > 0 || right < len(nums) - 1 {
        numLeft := 0
        if left > 0 {
            numLeft = nums[left - 1]
        }
        numRight := 0
        if right < len(nums) - 1 {
            numRight = nums[right + 1]
        }
        if numLeft < numRight {
            right++
            if nums[right] < currMin {
                currMin = nums[right]
            }
        } else {
            left--
            if nums[left] < currMin {
                currMin = nums[left]
            }
        }
        if v := currMin * (right - left + 1); v > res {
            res = v
        }
    }
    return res 
}