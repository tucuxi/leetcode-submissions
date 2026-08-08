func findMissingElements(nums []int) []int {
    set := make(map[int]bool)
    min := nums[0]
    max := nums[0]

    for _, n := range nums {
        set[n] = true
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }

    var res []int

    for i := min; i < max; i++ {
        if !set[i] {
            res = append(res, i)
        }
    }

    return res
}