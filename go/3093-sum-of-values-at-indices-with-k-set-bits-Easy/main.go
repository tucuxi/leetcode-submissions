func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i, n := range nums {
        bits := 0
        for i != 0 {
            bits++
            i &= i - 1
        }
        if bits == k {
            sum += n
        }
    }
    return sum
}