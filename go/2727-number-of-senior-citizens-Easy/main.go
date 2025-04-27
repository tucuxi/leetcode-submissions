func countSeniors(details []string) int {
    sixties := 0
    for _, d := range details {
        age := 10 * (d[11] - '0') + d[12] - '0'
        if age > 60 {
            sixties++
        }
    }
    return sixties
}