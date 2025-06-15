func lengthOfLIS(nums []int) int {
    a := make([]int, 0, len(nums))
    for _, n := range nums {
        if len(a) == 0 || n > a[len(a) - 1] {
            a = append(a, n)
        } else {
            k := sort.Search(len(a), func(i int) bool { return a[i] >= n })
            a[k] = n
        }
    }
    return len(a)
}