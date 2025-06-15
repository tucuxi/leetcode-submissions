func reorganizeString(s string) string {
    var (
        f [26]int
        h maxHeap
        res strings.Builder
        prev rune
    )
    for _, ch := range s {
        f[ch - 'a']++
    }
    for i, n := range f {
        if n > 0 {
            h = append(h, item{rune('a' + i), n})
        }
    }
    heap.Init(&h)
    for len(h) > 0 {
        u := heap.Pop(&h).(item)
        if u.value == prev {
            if len(h) == 0 {
                return ""
            }
            v := u
            u = heap.Pop(&h).(item)
            heap.Push(&h, v)
        }
        res.WriteRune(u.value)
        prev = u.value
        u.count--
        if u.count > 0 {
            heap.Push(&h, u)
        }
    }
    return res.String()
}

type item struct {
	value rune
	count int
}

type maxHeap []item

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(item)) }

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}