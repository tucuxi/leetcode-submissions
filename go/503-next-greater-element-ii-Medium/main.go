func nextGreaterElements(nums []int) []int {
    res := make([]int, len(nums))
    for i := range res {
        res[i] = -1
    }

    var stack []int

    popStack := func(n int) {
        for len(stack) > 0 {
            stackIndex := len(stack) - 1
            numIndex := stack[stackIndex]
            if nums[numIndex] >= n {
                break
            }
            res[numIndex] = n
            stack = stack[:stackIndex]
        }
    }

    for i, n := range nums {
        popStack(n)
        stack = append(stack, i)
    }
    
    for _, n := range nums {
        popStack(n)
    }

    return res
}