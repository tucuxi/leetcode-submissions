func intersection(nums1 []int, nums2 []int) []int {
    var res []int
    s1 := toSet(nums1)
    s2 := toSet(nums2)
    for n := range s1 {
        if s2[n] {
            res = append(res, n)
        }
    }
    return res
}

func toSet(nums []int) map[int]bool {
    res := make(map[int]bool)
    for _, n := range nums {
        res[n] = true
    }
    return res
}