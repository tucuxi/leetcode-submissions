func findLengthOfShortestSubarray(arr []int) int {
    right := len(arr) - 1
    for right > 0 && arr[right - 1] <= arr[right] {
        right--
    }
    res := right
    for left := 0; left < right && (left == 0 || arr[left - 1] <= arr[left]); left++ {
        for right < len(arr) && arr[left] > arr[right] {
            right++
        }
        res = min(res, right - left - 1)
    }
    return res
}
