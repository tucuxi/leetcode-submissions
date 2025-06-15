func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    var i int
    var s int64
    for j, n := range nums {
        s += int64(n)
        if int64((j - i + 1)) * int64(n) - s > int64(k) {
            s -= int64(nums[i])
            i++
        }
    }
    return len(nums) - i
}