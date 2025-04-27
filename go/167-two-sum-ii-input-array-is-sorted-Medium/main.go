func twoSum(numbers []int, target int) []int {
    p, q := 0, len(numbers) - 1
    for p < q {
        if sum := numbers[p] + numbers[q]; sum < target {
            p++
        } else if sum > target {
            q--
        } else {
            return []int{p + 1, q + 1}
        }
    }
    return nil
}