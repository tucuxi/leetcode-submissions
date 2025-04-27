func sumOfThree(num int64) []int64 {
    if num % 3 != 0 {
        return nil
    }
    k := num / 3
    return []int64{k - 1, k, k + 1}
}