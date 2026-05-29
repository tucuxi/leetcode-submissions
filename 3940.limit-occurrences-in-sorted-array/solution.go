func limitOccurrences(nums []int, k int) []int {
    res := []int{}
    p := 0
    c := 0

    for _, num := range nums {
        if num == p {
            c++
        } else {
            p = num
            c = 1
        }
        if c <= k {
            res = append(res, num)
        }
    }
    return res
}