func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    res := make([][]int, 0, k)
    pq := make(PriorityQueue, 0, k)
    heap.Push(&pq, &Item{0, 0, nums1[0] + nums2[0], -1})
    v := make(map[[2]int]bool)
    for k > 0 && len(pq) > 0 {
        next := heap.Pop(&pq).(*Item)
        if !v[[2]int{next.index1, next.index2}] {
            v[[2]int{next.index1, next.index2}] = true
            res = append(res, []int{nums1[next.index1], nums2[next.index2]})
            k--
            if next.index1 + 1 < len(nums1) {
                heap.Push(&pq, &Item{next.index1 + 1, next.index2, nums1[next.index1 + 1] + nums2[next.index2], -1})
            }
            if next.index2 + 1 < len(nums2) {
                heap.Push(&pq, &Item{next.index1, next.index2 + 1, nums1[next.index1] + nums2[next.index2 + 1], -1})
            }
        }
    }
    return res
}

type Item struct {
	index1   int
    index2   int    
	priority int    
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}