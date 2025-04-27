func removeDigit(number string, digit byte) string {
    p := 0
    for i := range number {
        if number[i] == digit {
            p = i
            if i < len(number) - 1 && number[i + 1] > digit {
                break
            }
        }
    }
    return number[:p] + number[p+1:]
}