func countTestedDevices(batteryPercentages []int) int {
    c, d := 0, 0
    for _, b := range batteryPercentages {
        if b > d {
            c++
            d++
        }
    }
    return c
}