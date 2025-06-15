func subarraysDivByK(nums []int, k int) int {
    modGroups := make([]int, k)
    modGroups[0] = 1
    prefixMod := 0
    res := 0
    for _, n := range nums {
        prefixMod = ((prefixMod + n) % k + k) % k
        res += modGroups[prefixMod]
        modGroups[prefixMod]++
    }
    return res
}