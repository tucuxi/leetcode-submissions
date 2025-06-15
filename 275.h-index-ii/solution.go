func hIndex(citations []int) int {
    return sort.Search(len(citations), func(i int) bool { return citations[len(citations) - 1 - i] <= i })
}
