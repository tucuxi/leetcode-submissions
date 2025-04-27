func checkIfExist(arr []int) bool {
    h := make(map[int]bool)
    for _, a := range arr {
        if h[2 * a] || a % 2 == 0 && h[a / 2] {
            return true
        }
        h[a] = true
    }
    return false
}