func answerQueries(nums []int, queries []int) []int {
    sort.Ints(nums)
    answer := make([]int, len(queries))
    for i := range queries {
        sum := 0
        j := 0
        for j < len(nums) && sum + nums[j] <= queries[i] {
            sum += nums[j]
            j++
        }
        answer[i] = j
    }
    return answer
}