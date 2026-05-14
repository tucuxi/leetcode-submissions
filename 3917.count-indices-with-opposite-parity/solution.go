func countOppositeParity(nums []int) []int {
    odd, even := 0, 0
    answer := make([]int, len(nums))

    for i := len(nums)-1; i >= 0; i-- {
        if nums[i]%2 == 0 {
            answer[i] = odd
            even++
        } else {
            answer[i] = even
            odd++
        }
    }

    return answer
}