func differenceOfSum(nums []int) int {
    elementSum := 0
    digitSum := 0
    for _, n := range nums {
        elementSum += n
        for k := n; k > 0; k /= 10 {
            digitSum += k % 10
        }
    }
    return abs(elementSum - digitSum)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}