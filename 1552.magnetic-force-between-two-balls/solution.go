func maxDistance(position []int, m int) int {
    sort.Ints(position)

    isValid := func(x int) bool {
        i := 0
        for range m - 1 {
            j := sort.Search(len(position) - i - 1, func(k int) bool {
                return position[i + 1 + k] >= position[i] + x
            })
            if (j == len(position) - i - 1) {
                return false
            }
            i += 1 + j
        }
        return true
    }

    lo := 1
    hi := (position[len(position) - 1] - position[0]) / (m - 1)
    for (lo < hi) {
        mid := lo + (hi - lo + 1) / 2
        if isValid(mid) {
            lo = mid
        } else {
            hi = mid - 1
        }
    }
    return lo
}