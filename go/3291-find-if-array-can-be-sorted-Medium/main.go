func canSortArray(nums []int) bool {
    var previousLargest, currentLargest uint
    currentBits := 0
    
    for _, num := range nums {
        n := uint(num)
        b := bits.OnesCount(n)
        if b == currentBits {
            currentLargest = max(currentLargest, n)
        } else {
            previousLargest = currentLargest
            currentLargest = n
            currentBits = b
        }
        if n < previousLargest {
            return false
        }
    }
    
    return true
}