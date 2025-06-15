func canReach(arr []int, start int) bool {
    v := make([]bool, len(arr))

    var visit func(int) bool
    visit = func(p int) bool {
        if p < 0 || p >= len(arr) || v[p] {
            return false
        }
        if arr[p] == 0 {
            return true
        }
        v[p] = true
        return visit(p - arr[p]) || visit(p + arr[p])
    }

    return visit(start)
}