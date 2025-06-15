func findIntersectionValues(nums1 []int, nums2 []int) []int {
    var count1, count2 [101]int
    
    for _, n1 := range nums1 {
        count1[n1]++
    }
    for _, n2 := range nums2 {
        count2[n2]++
    }
    
    res := make([]int, 2)
    
    for _, n1 := range nums1 {
        if count2[n1] > 0 {
            res[0]++
        }
    }
    for _, n2 := range nums2 {
        if count1[n2] > 0 {
            res[1]++
        }
    }
    
    return res
}