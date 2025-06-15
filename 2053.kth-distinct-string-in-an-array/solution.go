func kthDistinct(arr []string, k int) string {
    count := make(map[string]int)

    for _, s := range arr {
        count[s]++
    }

    j := 1

    for _, s := range arr {
        if count[s] == 1 {
            if j == k {
                return s
            }
            j++
        }
    }

    return ""
}