func nextGreaterElement(nums1 []int, nums2 []int) []int {
    return nextGreater(nums1, h(nums2))
}

func h(nums []int) map[int]int {
    res := make(map[int]int)
    st := make([]int, 0)
    for _, n := range nums {
        top := len(st) - 1
        for top >= 0 && st[top] < n {
            res[st[top]] = n
            st = st[:top]
            top--
        }
        st = append(st, n)
    }
    return res
}

func nextGreater(nums []int, h map[int]int) []int {
    res := make([]int, 0, len(nums))
    for _, m := range nums {
        if r, exists := h[m]; exists {
            res = append(res, r)
        } else {
            res = append(res, -1)
        }
    }
    return res
}