func minimumIndex(capacity []int, itemSize int) int {
    index := -1
    minCapacity := math.MaxInt
    for i, c := range capacity {
        if c >= itemSize && c < minCapacity {
            index = i
            minCapacity = c 
        }
    }
    return index
}