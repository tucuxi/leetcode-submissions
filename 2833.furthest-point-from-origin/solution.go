func furthestDistanceFromOrigin(moves string) int {
    left, right, any := 0, 0, 0

    for _, m := range moves {
        switch m {
            case 'L':
                left++
            case 'R':
                right++
            default:
                any++
        }
    }

    return max(right - left, left - right) + any
}