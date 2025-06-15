func maxHeightOfTriangle(red int, blue int) int {
    return max(height([2]int{red, blue}), height([2]int{blue, red}))
}

func height(balls [2]int) int {
    i := 0
    k := 1
    for balls[i] >= k {
        balls[i] -= k
        k++
        i = 1-i
    }
    return k-1
}