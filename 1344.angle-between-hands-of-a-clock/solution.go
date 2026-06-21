func angleClock(hour int, minutes int) float64 {
    h := float64(hour % 12 * 30) + float64(minutes) / 2
    m := float64(minutes) * 6
    a := math.Mod(h - m + 360, 360)
    return min(a, 360 - a)
}