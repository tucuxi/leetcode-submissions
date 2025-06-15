type entry struct {
    node int
    time int
}

const MOD = 1_000_000_007

func countPaths(n int, roads [][]int) int {
    graph := make([][]entry, n)
    for _, road := range roads {
        graph[road[0]] = append(graph[road[0]], entry{road[1], road[2]})
        graph[road[1]] = append(graph[road[1]], entry{road[0], road[2]})
    }

    time := make([]int, n)
    for i := range time {
        time[i] = math.MaxInt
    }
    time[0] = 0

    paths := make([]int, n)
    paths[0] = 1

    queue := make(entryHeap, 0, n)
    queue = append(queue, entry{0, 0})
    heap.Init(&queue)

    for len(queue) > 0 {
        u := heap.Pop(&queue).(entry)
        if u.time <= time[u.node] {
            for _, v := range graph[u.node] {
                if t := u.time + v.time; t < time[v.node] {
                    time[v.node] = t
                    paths[v.node] = paths[u.node]
                    heap.Push(&queue, entry{v.node, t})
                } else if t == time[v.node] {
                    paths[v.node] = (paths[v.node] + paths[u.node]) % MOD
                }
            }
        }
    }
    return paths[n - 1]
}

type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *entryHeap) Push(x any)        { *h = append(*h, x.(entry)) }

func (h *entryHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}