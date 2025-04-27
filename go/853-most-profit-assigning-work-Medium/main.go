func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    index := make([]int, len(difficulty))
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(a, b int) int { return difficulty[a] - difficulty[b] })
    slices.Sort(worker)
    maxProfitTotal := 0
    maxProfitJob := 0
    mostDifficultJob := 0
    for _, w := range worker {
        for mostDifficultJob < len(index) && difficulty[index[mostDifficultJob]] <= w {
            maxProfitJob = max(maxProfitJob, profit[index[mostDifficultJob]])
            mostDifficultJob++ 
        }
        maxProfitTotal += maxProfitJob
    }
    return maxProfitTotal
}