func numberOfSubarrays(nums []int, k int) int {
    subarrays := 0
    lastPopped := -1
    initialGap := -1
    var oddIndices []int
    
    for i, num := range nums {
        if num % 2 == 1 {
            oddIndices = append(oddIndices, i)
        }
        if len(oddIndices) > k {
            lastPopped = oddIndices[0]
            oddIndices = oddIndices[1:]
        }
        if len(oddIndices) == k {
            initialGap = oddIndices[0] - lastPopped
            subarrays += initialGap
        }
    }

    return subarrays
}