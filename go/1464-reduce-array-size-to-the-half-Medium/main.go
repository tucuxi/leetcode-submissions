func minSetSize(arr []int) int {
    h := make(map[int]int) 
    for _, a := range arr {
        h[a]++
    }
    var counts []int
    for _, n := range h {
        counts = append(counts, n)
    }
    sort.Ints(counts)
    i := len(counts)
    for n := len(arr); n > len(arr) / 2; n -= counts[i] {
        i--
    }
    return len(counts) - i
}