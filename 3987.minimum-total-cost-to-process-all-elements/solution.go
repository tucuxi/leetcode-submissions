func minimumCost(nums []int, k int) int {
    const mod = 1_000_000_007
    sum := 0

    for _, num := range nums {
        sum += num
    }

    ops := (sum + k - 1) / k % mod
    return (ops - 1) * ops / 2 % mod   
}
