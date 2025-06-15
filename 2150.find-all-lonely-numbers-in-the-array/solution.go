func findLonely(nums []int) []int {
    h := make(map[int]int)
    for _, num := range nums {
        h[num]++
    }

    var res []int
    for num := range h {
        if h[num] == 1 && h[num - 1] == 0 && h [num + 1] == 0 {
            res = append(res, num)
        }
    }

    return res
}