func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1) + len(nums2)
    if n % 2 != 0 {
        return solve(nums1, nums2, n / 2, 0, len(nums1) - 1, 0, len(nums2) - 1)
    }
    return (solve(nums1, nums2, n / 2 - 1, 0, len(nums1) - 1, 0, len(nums2) - 1) +
            solve(nums1, nums2, n / 2, 0, len(nums1) - 1, 0, len(nums2) - 1)) / 2.0
}

func solve(a, b []int, k, aStart, aEnd, bStart, bEnd int) float64 {
    if aEnd < aStart {
        return float64(b[k - aStart])
    }
    if bEnd < bStart {
        return float64(a[k - bStart])
    }
    aIndex := (aStart + aEnd) / 2
    aValue := a[aIndex]
    bIndex := (bStart + bEnd) / 2
    bValue := b[bIndex]
    if aIndex + bIndex < k {
        if aValue > bValue {
            return solve(a, b, k, aStart, aEnd, bIndex + 1, bEnd)
        }
        return solve(a, b, k, aIndex + 1, aEnd, bStart, bEnd)
    } else {
        if aValue > bValue {
            return solve(a, b, k, aStart, aIndex - 1, bStart, bEnd)
        }
        return solve(a, b, k, aStart, aEnd, bStart, bIndex - 1)
    }
    return -1
}