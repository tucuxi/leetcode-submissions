func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    if numOnes >= k {
        return k
    }
    if numOnes + numZeros >= k {
        return numOnes
    }
    return 2 * numOnes + numZeros - k
}