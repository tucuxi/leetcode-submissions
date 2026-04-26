func solveQueries(nums []int, queries []int) []int {
    h := make(map[int][]int)
    
    for i, n := range nums {
        h[n] = append(h[n], i)
    }

    answer := make([]int, 0, len(queries))
    
    for _, q := range queries {
        indices := h[nums[q]]
        a := -1
        if len(indices) > 1 {
            i := sort.Search(len(indices), func(i int) bool { return indices[i] >= q })
            switch i {
            case 0:
                a = min(indices[i] + len(nums) - indices[len(indices) - 1], indices[i+1] - indices[i])
            case len(indices) - 1:
                a = min(indices[i] - indices[i-1], len(nums) - indices[i] + indices[0])
            default:
                a = min(indices[i] - indices[i-1], indices[i+1] - indices[i])
            }
        }
        answer = append(answer, a)
    }

    return answer
}