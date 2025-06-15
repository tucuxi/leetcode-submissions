func maxDistance(colors []int) int {
    dist := 0
    for i := 0; i < len(colors); i++ {
        for j := i + 1; j < len(colors); j++ {
            if colors[i] != colors[j] && j - i > dist {
                dist = j - i
            }
        }
    }
    return dist
}