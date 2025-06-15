func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i, j int) bool {
        p1, p2 := people[i], people[j]
        return p1[0] > p2[0] || p1[0] == p2[0] && p1[1] < p2[1]
    })
    return reconstructFromOrderedQueue(people)
}

func reconstructFromOrderedQueue(people [][]int) [][]int {
    ans := make([][]int, len(people))
    k := 0
    for _, p := range people {
        i := p[1]
        for j := k; j > i; j-- {
            ans[j] = ans[j - 1]
        }
        ans[i] = p
        k++
    }
    return ans
}