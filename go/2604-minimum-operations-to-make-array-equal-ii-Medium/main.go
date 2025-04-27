func minOperations(nums1 []int, nums2 []int, k int) int64 {
    equal := true
    var diffP, diffN int64
    for i := range nums1 {
        if diff := nums1[i] - nums2[i]; diff > 0 {
            diffP += int64(diff)
            equal = false
        } else if diff < 0 {
            diffN += int64(-diff)
            equal = false
        }
    }
    if k == 0 {
        if equal {
            return 0
        }
        return -1
    }
    if diffP == diffN && diffP % int64(k) == 0 {
        return diffP / int64(k)
    }
    return -1
}