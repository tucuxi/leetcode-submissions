func minimumSize(nums []int, maxOperations int) int {
    i := 1
    j := 0

    for _, n := range nums {
        j = max(j, n)
    }

    for i < j {
        k := (i + j) / 2
        if operations(k, nums) <= maxOperations {
            j = k
        } else {
            i = k + 1
        }
    }

    return i
}

func operations(maxBallsInBag int, nums []int) int {
    res := 0
    for _, n := range nums {
        res += (n + maxBallsInBag - 1) / maxBallsInBag - 1
    }
    return res
}