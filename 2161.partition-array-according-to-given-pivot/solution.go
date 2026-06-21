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
    i, j, k := 0, less, less + equal
    for _, n := range nums {
        if n < pivot {
            res[i] = n
            i++
        } else if n == pivot {
            res[j] = n
            j++
        } else {
            res[k] = n
            k++
        }
    }

    return res
}