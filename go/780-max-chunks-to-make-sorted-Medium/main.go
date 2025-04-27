func maxChunksToSorted(arr []int) int {
    chunks := 0
    maxSeen := 0

    for i, a := range arr {
        maxSeen = max(maxSeen, a)
        if i == maxSeen {
            chunks++
        }
    }

    return chunks
}