func validMountainArray(arr []int) bool {
    i := 0
    for i < len(arr) - 1 && arr[i] < arr[i + 1] {
        i++
    }
    if i == 0 || i + 1 == len(arr) {
        return false
    }
    for i < len(arr) - 1 && arr[i] > arr[i + 1] {
        i++
    }
    return i + 1 == len(arr)
}