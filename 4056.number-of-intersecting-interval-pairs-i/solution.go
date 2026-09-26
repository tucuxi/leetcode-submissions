func countIntersectingIntervals(intervals [][]int) int {
    res := 0
    for i := range intervals {
        s1, e1 := intervals[i][0], intervals[i][1]
        for j := i+1; j < len(intervals); j++ {
            s2, e2 := intervals[j][0], intervals[j][1]
            if s1 <= s2 && e1 >= s2 || s2 <= s1 && e2 >= s1 {
                res++
            }  
        }
    }
    return res
}