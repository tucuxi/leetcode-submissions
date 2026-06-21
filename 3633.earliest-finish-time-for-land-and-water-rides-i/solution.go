func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    t1 := optimize(landStartTime, landDuration, waterStartTime, waterDuration)
    t2 := optimize(waterStartTime, waterDuration, landStartTime, landDuration)
    return min(t1, t2)
}

func optimize(s1, d1, s2, d2 []int) int {
    t1, t2 := math.MaxInt, math.MaxInt

    for i := range s1 {
        t1 = min(s1[i] + d1[i], t1)
    }
    for i := range s2 {
        t2 = min(max(s2[i], t1) + d2[i], t2)
    }

    return t2
}