func minCost(colors string, neededTime []int) int {
    curMaxTime, totalTime := 0, 0
    for i := range colors {
        if i > 0 && colors[i - 1] != colors[i] {
            curMaxTime = 0
        }
        if neededTime[i] > curMaxTime {
            totalTime += curMaxTime
            curMaxTime = neededTime[i]
        } else {
            totalTime += neededTime[i]
        }
    }
    return totalTime
}