func semiOrderedPermutation(nums []int) int {
    n := len(nums)
    var index1, indexn int
    for i, k := range nums {
        if k == 1 {
            index1 = i
        }
        if k == n {
            indexn = i
        }
    }
    if index1 > indexn {
        return index1 + n - 2 - indexn
    }
    return index1 + n - 1 - indexn
}