func maxArea(height []int) int {
    volume := 0
    left, right := 0, len(height) - 1
    for left < right {
        volume = max(volume, min(height[left], height[right]) * (right - left))
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return volume
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}