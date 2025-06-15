type entry struct {
    y int
    x int
}

func orangesRotting(grid [][]int) int {
    oranges := 0
    queue := []entry{}
    for y := range grid {
        for x := range grid[y] {
            switch grid[y][x] {
                case 1: oranges++
                case 2: queue = append(queue, entry{y, x})
            }
        }
    }
    t := 0
    for len(queue) > 0 {
        for i := len(queue); i > 0; i-- {
            y, x := queue[0].y, queue[0].x
            queue = queue[1:]
            grid[y][x] = 0
            for _, nb := range []entry{{y - 1, x}, {y + 1, x}, {y, x - 1}, {y, x + 1}} {
                if nb.y >= 0 && nb.y < len(grid) && nb.x >= 0 && nb.x < len(grid[nb.y]) {
                    if grid[nb.y][nb.x] == 1 {
                        grid[nb.y][nb.x] = 2
                        queue = append(queue, nb)
                        oranges--
                    }
                }
            }
        }
        t++
    }
    if oranges > 0 {
        return -1
    }
    if t == 0 {
        return 0
    }
    return t - 1
}