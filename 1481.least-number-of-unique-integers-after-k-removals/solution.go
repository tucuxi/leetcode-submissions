func findLeastNumOfUniqueInts(arr []int, k int) int {
    freq := frequencies(arr)
    i := 0
    for i < len(freq) && freq[i] <= k {
        k -= freq[i]
        i++
    }
    return len(freq) - i
}

func frequencies(arr []int) []int {
    h := make(map[int]int)
    for _, n := range arr {
        h[n]++
    }
    res := make([]int, 0, len(h))
    for _, v := range h {
        res = append(res, v)
    }
    sort.Ints(res)
    return res
}