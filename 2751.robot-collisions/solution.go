func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    index := make([]int, len(positions))

    for i := range index {
        index[i] = i
    }

    slices.SortFunc(index, func(a, b int) int { return positions[a] - positions[b] })

    var rightMovers []int

    for _, robot := range index {
        if directions[robot] == 'R' {
            rightMovers = append(rightMovers, robot)
        } else {
            for len(rightMovers) > 0 && healths[robot] > 0 {
                j := len(rightMovers) - 1
                collidingRobot := rightMovers[j]
                rightMovers = rightMovers[:j]
                if healths[robot] < healths[collidingRobot] {
                    healths[robot] = 0
                    healths[collidingRobot]--
                    rightMovers = append(rightMovers, collidingRobot)
                } else if healths[robot] > healths[collidingRobot] {
                    healths[robot]--
                    healths[collidingRobot] = 0
                } else {
                    healths[robot] = 0
                    healths[collidingRobot] = 0
                }
            }
        }
    }

    var res []int

    for i := range healths {
        if healths[i] > 0 {
            res = append(res, healths[i])
        }
    }

    return res
}