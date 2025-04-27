func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
    positiveSet, negativeSet := map[string]bool{}, map[string]bool{}
    for _, positive := range positive_feedback {
        positiveSet[positive] = true
    }
    for _, negative := range negative_feedback {
        negativeSet[negative] = true
    }
    points := [][2]int{}
    for i := range report {
        score := 0
        words := strings.Split(report[i], " ")
        for _, w := range words {
            if positiveSet[w] {
                score += 3
            } else if negativeSet[w] {
                score--
            }
        }
        points = append(points, [2]int{student_id[i], score})        
    }
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] > points[j][1] || points[i][1] == points[j][1] && points[i][0] < points[j][0]
    })
    res := make([]int, k)
    for i := range res {
        res[i] = points[i][0]
    }
    return res
}