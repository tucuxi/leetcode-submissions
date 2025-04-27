func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    originalColor := image[sr][sc]
    if color == originalColor {
        return image
    }
    for queue := [][2]int{{sr, sc}}; len(queue) > 0; queue = queue[1:] {
        row, col := queue[0][0], queue[0][1]
        image[row][col] = color
        for _, neighbor := range [][2]int{{row - 1, col}, {row + 1, col}, {row, col -1}, {row, col + 1}} {
            if neighbor[0] < 0 || neighbor[0] == len(image) || neighbor[1] < 0 || neighbor[1] == len(image[neighbor[0]]) {
                continue
            }
            if image[neighbor[0]][neighbor[1]] != originalColor {
                continue
            }
            queue = append(queue, neighbor)
        }
    }
    return image
}