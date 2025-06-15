type entry struct {
    employee int
    time int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    sub := make([][]int, n)
    for i, m := range manager {
        if m != -1 {
            sub[m] = append(sub[m], i)
        }
    }
    res := 0
    queue := []entry{{headID, 0}}
    for len(queue) > 0 {
        m := queue[0].employee
        t := queue[0].time + informTime[m]
        queue = queue[1:]
        for _, e := range sub[m] {
            queue = append(queue, entry{e, t})
        }
        if t > res {
            res = t
        }
    }
    return res
}