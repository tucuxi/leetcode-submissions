type entry struct {
    row, col, time, turn int
}

func minTimeToReach(moveTime [][]int) int {
    m, n := len(moveTime), len(moveTime[0])

    time := make([][]int, m)
    for i := range m {
        time[i] = make([]int, n)
        for j := range n {
            time[i][j] = math.MaxInt
        }
    }
    time[0][0] = 0

    q := entryHeap{entry{0, 0, 0, 1}}

    improve := func(dr, dc int, e entry) {
        r, c := e.row + dr, e.col + dc
        newTime := max(e.time, moveTime[r][c]) + e.turn
        if newTime < time[r][c] {
            time[r][c] = newTime
            heap.Push(&q, entry{r, c, newTime, 3 - e.turn})
        }
    }

    for q.Len() > 0 {
        e := heap.Pop(&q).(entry)
        if e.row == m - 1 && e.col == n - 1 {
            break
        }
        if e.row > 0 {
            improve(-1, 0, e)
        }
        if e.col > 0 {
            improve(0, -1, e)
        }
        if e.col + 1 < n {
            improve(0, 1, e)
        }
        if e.row + 1 < m {
            improve(1, 0, e)
        }
    }

    return time[m - 1][n - 1]
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