func findWinners(matches [][]int) [][]int {
    losses := map[int]int{}
    for _, m := range matches {
        if _, ok := losses[m[0]]; !ok {
            losses[m[0]] = 0
        }
        losses[m[1]]++
    } 
    answer := [][]int{[]int{}, []int{}}
    for k, v := range losses {
        if v == 0 {
            answer[0] = append(answer[0], k)
        } else if v == 1 {
            answer[1] = append(answer[1], k)
        }
    }
    sort.Ints(answer[0])
    sort.Ints(answer[1])
    return answer
}