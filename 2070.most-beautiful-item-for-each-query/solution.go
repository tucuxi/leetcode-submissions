func maximumBeauty(items [][]int, queries []int) []int {
    slices.SortFunc(items, func(a, b []int) int {
        d := a[0] - b[0]
        if d == 0 {
            return b[1] - a[1]
        }
        return d
    })

    table := make([][2]int, 0, len(items))
    maxBeauty := 0
    maxPrice := 0
    for _, item := range items {
        if item[0] > maxPrice {
            maxPrice = item[0]
            maxBeauty = max(maxBeauty, item[1])
            table = append(table, [2]int{maxPrice, maxBeauty})            
        }
    } 

    answer := make([]int, 0, len(queries))
    for _, query := range queries {
        index, found := slices.BinarySearchFunc(table, query, func(elem [2]int, price int) int {
            return elem[0] - price
        })
        if found {
            answer = append(answer, table[index][1])
        } else if index == 0 {
            answer = append(answer, 0)
        } else {
            answer = append(answer, table[index - 1][1])
        }
    }

    return answer
}