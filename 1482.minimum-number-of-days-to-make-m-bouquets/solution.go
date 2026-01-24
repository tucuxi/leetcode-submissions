func minDays(bloomDay []int, m int, k int) int {
    lastDay := slices.Max(bloomDay)

    day := sort.Search(lastDay + 1, func(i int) bool {
        flowers := 0
        bouquets := 0

        for _, day := range bloomDay {
            if day > i {
                flowers = 0
            } else {
                flowers++
                if flowers == k {
                    flowers = 0
                    bouquets++
                    if bouquets == m {
                        return true
                    }
                }
            }
        }

        return false
    })

    if day > lastDay {
        return -1
    }
    return day
}