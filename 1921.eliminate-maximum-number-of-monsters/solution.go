func eliminateMaximum(dist []int, speed []int) int {
    eta := make([]int, len(dist))
    for i := range eta {
        eta[i] = (dist[i] + speed[i] - 1) / speed[i] 
    }
    sort.Slice(eta, func(i, j int) bool { return eta[i] < eta[j] })
    i := 0
    for i < len(eta) && eta[i] > i {
        i++
    }
    return i
}