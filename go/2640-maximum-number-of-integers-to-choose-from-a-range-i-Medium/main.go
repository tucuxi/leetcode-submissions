func maxCount(banned []int, n int, maxSum int) int {
    sort.Ints(banned)

    sum := 0
    count := 0

    outer:
    for i := 1; i <= n && sum + i <= maxSum; i++ {
        if len(banned) > 0 && i == banned[0] {
            for {
                banned = banned[1:]
                if len(banned) == 0 || banned[0] > i {
                    continue outer
                }
            }
        }
        sum += i
        count++
    }

    return count
}