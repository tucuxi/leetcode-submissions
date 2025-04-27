func averageWaitingTime(customers [][]int) float64 {
    w := 0
    t := 0
    for _, c := range customers {     
        t = max(t, c[0]) + c[1]   
        w += t - c[0]

    }
    return float64(w) / float64(len(customers))
}