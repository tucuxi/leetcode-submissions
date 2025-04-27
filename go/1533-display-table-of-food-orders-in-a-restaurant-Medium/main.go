func displayTable(orders [][]string) [][]string {
    ordersPerTable, allItems := func() (map[int]map[string]int, map[string]bool) {
        ordersPerTable := make(map[int]map[string]int)        
        allItems := make(map[string]bool)
        for _, o := range orders {
            table, food := o[1], o[2]
            tableNum, _ := strconv.Atoi(table)
            allItems[food] = true
            f, exists := ordersPerTable[tableNum]
            if !exists {
                f = make(map[string]int)
                ordersPerTable[tableNum] = f
            }
            f[food]++
        }
        return ordersPerTable, allItems
    }()
    sortedTables := func() []int {
        tt := make([]int, 0, len(ordersPerTable))
        for t := range ordersPerTable {
            tt = append(tt, t)
        }
        sort.Ints(tt)
        return tt
    }()
    sortedItems := func() []string {
        ii := make([]string, 0, len(allItems))
        for i := range allItems {
            ii = append(ii, i)
        }
        sort.Strings(ii)
        return ii
    }()
    res := func() [][]string {
        res := make([][]string, len(sortedTables) + 1)
        res[0] = make([]string, 1, len(sortedItems) + 1)
        res[0][0] = "Table"
        res[0] = append(res[0], sortedItems...)
        k := 0
        for _, t := range sortedTables {
            k++
            res[k] = make([]string, 1, len(sortedItems) + 1)
            res[k][0] = strconv.Itoa(t)
            for _, i := range sortedItems {
                res[k] = append(res[k], strconv.Itoa(ordersPerTable[t][i]))
            }
        }
        return res
    }()
    return res
}