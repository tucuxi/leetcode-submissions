func minimumTime(grid [][]int) int {
    if grid[0][1] > 1 && grid[1][0] > 1 {
        return -1
    }
    
    m, n := len(grid), len(grid[0])

    minTime := make([][]int, m)
    for i := range m {
        minTime[i] = make([]int, n)
        for j := range n {
            minTime[i][j] = math.MaxInt
        }
    }
    minTime[0][0] = 0

    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1 , 0}}

    h := entryHeap{{0, 0, 0}}

    for {
        current := heap.Pop(&h).(entry)
        if current.time > minTime[current.row][current.col] {
            continue
        }
        if current.row == m - 1 && current.col == n - 1 {
            return current.time
        }
        for _, dir := range directions {
            row, col := current.row + dir[0], current.col + dir[1]
            if row < 0 || row >= m || col < 0 || col >= n {
                continue
            }
            time := current.time + 1
            if time < grid[row][col] {
                time = grid[row][col] + (grid[row][col] - time) % 2
            }
            if time < minTime[row][col] {
                minTime[row][col] = time
                heap.Push(&h, entry{time, row, col})
            }
        }    
    }
}

type entry struct {
    time, row, col int
}

type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *entryHeap) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *entryHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}