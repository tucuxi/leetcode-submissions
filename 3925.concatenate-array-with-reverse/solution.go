func concatWithReverse(nums []int) []int {
    ans := make([]int, 2*len(nums))
    j := len(ans)
    for i, num := range nums {
        ans[i] = num
        j--
        ans[j] = num
    }
    return ans
}