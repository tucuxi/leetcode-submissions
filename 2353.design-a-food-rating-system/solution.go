type FoodRatings struct {
    f map[string]*Item
    r map[string]*PriorityQueue
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    f := make(map[string]*Item)
    r := make(map[string]*PriorityQueue)
    for i := range foods {
        item := &Item{
            food:    foods[i],
            cuisine: cuisines[i],
            rating:  ratings[i],
            index:   -1,
        }
        f[foods[i]] = item
        if r[cuisines[i]] == nil {
            r[cuisines[i]] = new(PriorityQueue)
        }
        heap.Push(r[cuisines[i]], item)
    }
    return FoodRatings{f, r}
}


func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    item := this.f[food]
    item.rating = newRating
    heap.Fix(this.r[item.cuisine], item.index)
}


func (this *FoodRatings) HighestRated(cuisine string) string {
    return (*this.r[cuisine])[0].food
}

type Item struct {
	food     string 
    cuisine  string
	rating   int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].rating > pq[j].rating || pq[i].rating == pq[j].rating && pq[i].food < pq[j].food
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
	old[n-1] = nil 
	item.index = -1
	*pq = old[0 : n-1]
	return item
}