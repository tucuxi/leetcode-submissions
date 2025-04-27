const maxQuantity = 100000

func minimizedMaximum(n int, quantities []int) int {
    canBeDistributed := func(x int) bool {
        shopsNeeded := 0
        for _, q := range quantities {
            shopsNeeded += (q + (x - 1)) / x 
        }
        return shopsNeeded <= n
    }
    left := 1
    right := maxQuantity
    for left < right {
        mid := left + (right - left) / 2
        if canBeDistributed(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return right
}