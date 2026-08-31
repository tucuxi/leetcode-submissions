func minBishopMoves(source []int, target []int) int {
    if source[0] == target[0] && source[1] == target[1] {
        return 0
    }
    if abs(source[0] - target[0]) == abs(source[1]- target[1]) {
        return 1
    }
    if (source[0] + source[1]) % 2 == (target[0] + target[1]) % 2 {
        return 2
    }
    return -1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}