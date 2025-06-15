const vmax = 1_000_000_001

func minSpeedOnTime(dist []int, hour float64) int {
    v := sort.Search(vmax, func(i int) bool {
        r := len(dist) - 1
        v := float64(i)
        t := 0.
        for _, d := range dist[:r] {
            t += math.Ceil(float64(d) / v) 
        }
        t += float64(dist[r]) / v
        return t <= hour
    })
    if v == vmax {
        return -1
    }
    return v
}