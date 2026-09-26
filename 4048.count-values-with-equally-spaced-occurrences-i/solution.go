func countSpecialIntegers(nums []int) int {
    h := make([][]int, 101)    
    for i, num := range nums {
        h[num] = append(h[num], i)
    }

    res := 0
    for _, indices := range h {
        if len(indices) == 3 && indices[1] - indices[0] == indices[2] - indices[1] {
            res++
        }
    }

    return res
}