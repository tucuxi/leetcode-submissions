/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := make(minHeap, 0, len(lists))
    for _, list := range lists {
        if list != nil {
            heap.Push(&h, list)
        }
    }
    sentinel := new(ListNode)
    p := sentinel
    for h.Len() > 0 {
        p.Next = heap.Pop(&h).(*ListNode)
        p = p.Next
        if p.Next != nil {
            heap.Push(&h, p.Next)
        }
    }
    return sentinel.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}