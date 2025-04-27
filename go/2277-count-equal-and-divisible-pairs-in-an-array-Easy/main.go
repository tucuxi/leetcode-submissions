func countPairs(nums []int, k int) int {
    h := make(map[int][]int)
    for i, num := range nums {
        h[num] = append(h[num], i)
    }

    res := 0
    for i, num := range nums {
        h[num] = h[num][1:]
        for _, j := range h[num] {
            if i * j % k == 0 {
                res++
            }
        }
    }
    return res
}