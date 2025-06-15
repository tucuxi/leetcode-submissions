func pivotArray(nums []int, pivot int) []int {
    less, equal := 0, 0
    for _, n := range nums {
        if n < pivot {
            less++
        } else if n == pivot {
            equal++
        }
    }

    res := make([]int, len(nums))
    p, q, r := 0, less, less + equal
    for _, n := range nums {
        if n < pivot {
            res[p] = n
            p++
        } else if n == pivot {
            res[q] = n
            q++
        } else {
            res[r] = n
            r++
        }
    }

    return res
}