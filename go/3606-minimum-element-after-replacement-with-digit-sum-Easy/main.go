func minElement(nums []int) int {
    res := math.MaxInt
    for _, num := range nums {
        sum := 0
        for num > 0 {
            sum += num % 10
            num /= 10
        }
        if sum < res {
            res = sum
        }
    }    
    return res
}