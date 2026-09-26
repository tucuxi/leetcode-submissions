func smallestIndex(nums []int) int {
    for i, n := range nums {
        if i == digitsSum(n) {
            return i
        }
    }
    return -1
}

func digitsSum(n int) int {
    sum := 0
    for ; n > 0; n /= 10 {
        sum += n%10
    }
    return sum
}