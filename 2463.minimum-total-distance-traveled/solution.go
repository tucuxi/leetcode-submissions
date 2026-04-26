func minimumTotalDistance(robot []int, factory [][]int) int64 {
    slices.Sort(robot)
    slices.SortFunc(factory, func(a, b []int) int { return a[0]-b[0] })

    var factoryPositions []int
    
    for _, f := range factory {
        for range f[1] {
            factoryPositions = append(factoryPositions, f[0])
        }
    }

    robotCount := len(robot)
    factoryCount := len(factoryPositions)
    dp := make([][]int64, robotCount+1)

    for i := range robotCount {
        dp[i] = make([]int64, factoryCount+1)
        dp[i][factoryCount] = 1_000_000_000_000
    }
    dp[robotCount] = make([]int64, factoryCount+1)

    for i := robotCount-1; i >= 0; i-- {
        for j := factoryCount-1; j >= 0; j-- {
            assign := int64(abs(robot[i] - factoryPositions[j])) + dp[i+1][j+1]
            skip := dp[i][j+1]
            dp[i][j] = min(assign, skip)
        }
    }

    return dp[0][0]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}