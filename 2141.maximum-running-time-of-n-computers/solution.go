func maxRunTime(n int, batteries []int) int64 {
    var power int64
    for _, b := range batteries {
        power += int64(b)
    }
    var lo, hi int64 = 1, power / int64(n)
    for lo < hi {
        target := hi - (hi - lo) / 2
        var extra int64
        for _, b := range batteries {
            extra += min(int64(b), target)
        }
        if extra >= int64(n) * target {
            lo = target
        } else {
            hi = target - 1
        }
    }
    return lo
}