type entry struct{ city, price int }

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    graph := make([][][]int, n) 
    for _, f := range flights {
        graph[f[0]] = append(graph[f[0]], f)
    }
    price := make([]int, n)
    for i := range price {
        price[i] = math.MaxInt
    }
    queue := []entry{{src, 0}}
    for i := 0; i <= k; i++ {
        for j := len(queue); j > 0; j-- {
            curr := queue[0]
            queue = queue[1:]
            for _, f := range graph[curr.city] {
                next := f[1]
                if p := curr.price + f[2]; p < price[next] {
                    price[next] = p
                    queue = append(queue, entry{next, p})
                }
            }
        }
    } 
    if price[dst] == math.MaxInt {
        return -1
    }
    return price[dst]
}