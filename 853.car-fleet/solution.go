type vehicle struct {
    position int
    speed    int
}

func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    v := make([]vehicle, n)

    for i := range n {
        v[i].position = position[i]
        v[i].speed = speed[i]
    }

    slices.SortFunc(v, func(a, b vehicle) int { return a.position - b.position })

    t := 0.0
    ft := float64(target)
    res := 0

    for i := n-1; i >= 0; i-- {
        if float64(v[i].position) + t * float64(v[i].speed) < ft {
            t = (ft - float64(v[i].position)) / float64(v[i].speed)
            res++
        }
    }

    return res
}