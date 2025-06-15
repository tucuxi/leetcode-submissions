func prefixesDivBy5(nums []int) []bool {
    answer := make([]bool, 0, len(nums))
    prefix := 0
    for _, n := range nums {
        prefix = (2 * prefix + n) % 10
        answer = append(answer, prefix == 0 || prefix == 5)
    }
    return answer
}