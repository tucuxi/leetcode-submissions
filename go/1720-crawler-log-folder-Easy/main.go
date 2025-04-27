func minOperations(logs []string) int {
    k := 0
    for _, op := range logs {
        if op == "../" {
            if k > 0 {
                k--
            }
        } else if op != "./" {
            k++
        }
    }
    return k
}