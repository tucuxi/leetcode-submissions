func stringSequence(target string) []string {
    var res []string
    for i := range target {
        for ch := byte('a'); ch <= target[i]; ch++ {
            res = append(res, target[:i] + string(ch))
        }
    }
    return res
}