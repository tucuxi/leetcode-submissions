func nearestDrone(drones [][]int, target []int) int {
    manhattan := func(p []int) int {
        return abs(p[0] - target[0]) + abs(p[1] - target[1])
    }

    minDist := math.MaxInt
    minIndex := -1

    for i, d := range drones {
        dist := manhattan(d)
        if dist <= d[2] && dist < minDist {
            minDist = dist
            minIndex = i
        }
    }

    return minIndex
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
