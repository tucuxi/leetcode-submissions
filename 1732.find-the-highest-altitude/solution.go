func largestAltitude(gain []int) int {
    altitude, highest := 0, 0
    for _, g := range gain {
        altitude += g
        if altitude > highest {
            highest = altitude
        }
    }
    return highest
}