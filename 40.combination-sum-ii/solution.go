func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int

    var rec func(int, int, []int)
    rec = func(i int, sum int, combination []int) {
        if sum == target {
            solution := make([]int, len(combination))
            copy(solution, combination)
            res = append(res, solution)
            return
        }
        if sum > target || i == len(candidates) {
            return
        }
        for j := i; j < len(candidates); j++ {
            if j == i || candidates[j] != candidates[j - 1] {
                rec(j + 1, sum + candidates[j], append(combination, candidates[j]))
            }
        }
    }

    slices.Sort(candidates)
    rec(0, 0, make([]int, 0, len(candidates)))
    return res
}