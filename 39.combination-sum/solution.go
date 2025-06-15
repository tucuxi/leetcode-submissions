func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}

    var rec func([]int, int, int)
    rec = func(pre []int, sum int, start int) {
        if sum > target {
            return
        }
        if sum == target {
            solution := make([]int, len(pre))
            copy(solution, pre)
            res = append(res, solution)
            return
        }
        for i := start; i < len(candidates); i++ {
            rec(append(pre, candidates[i]), sum + candidates[i], i)
        }
    }

    sort.Ints(candidates)
    rec(make([]int, 0, len(candidates)), 0, 0)
    return res
}