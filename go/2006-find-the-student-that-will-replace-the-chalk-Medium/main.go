func chalkReplacer(chalk []int, k int) int {
    sum := make([]int, len(chalk))
    pre := 0
    for i, c := range chalk {
        pre += c
        sum[i] = pre
    }
    student, _ := slices.BinarySearch(sum, k % pre + 1)
    return student
}