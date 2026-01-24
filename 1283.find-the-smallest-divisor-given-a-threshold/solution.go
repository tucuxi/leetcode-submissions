func smallestDivisor(nums []int, threshold int) int {
    return 1 + sort.Search(1000000, func(i int) bool {
        sum := 0
        for _, num := range nums {
            sum += (num + i) / (i + 1)
        }
        return sum <= threshold
    })
}