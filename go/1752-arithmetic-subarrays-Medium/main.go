func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    res := make([]bool, len(l))
    tmp := make([]int, len(nums))
    
    outer:
    for i := range l {
        left := l[i]
        right := r[i]
        copy(tmp, nums[left:right + 1])
        sort.Ints(tmp[:right -left + 1])
        d := tmp[1] - tmp[0]
        for j := 2; j <= right - left; j++ {
            if tmp[j] - tmp[j - 1] != d {
                continue outer
            }
        }
        res[i] = true
    }
    
    return res
}