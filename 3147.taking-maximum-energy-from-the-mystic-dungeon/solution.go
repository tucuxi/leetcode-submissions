func maximumEnergy(energy []int, k int) int {
    runningEnergy := make([]int, k)
    maxEnergy := math.MinInt
    for i := len(energy) - 1; i >= 0; i-- {
        j := i % k
        runningEnergy[j] += energy[i]
        maxEnergy = max(runningEnergy[j], maxEnergy)
    }
    return maxEnergy
}