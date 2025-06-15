func findKthPositive(arr []int, k int) int {
    return k + sort.Search(len(arr), func(i int) bool { return arr[i] > k + i })
}