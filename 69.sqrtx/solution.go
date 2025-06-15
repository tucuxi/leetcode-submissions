func mySqrt(x int) int {
    left := 1
    right := x
    for left < right {
        mid := left + (right - left) / 2
        sq := mid * mid
        if sq == x {
            return mid
        }
        if sq > x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left * left > x {
        return left - 1
    }
    return left
}