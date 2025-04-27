func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    res := make([]int, 0, len(obstacles))
    lis := make([]int, 0, len(obstacles))
    for _, o := range obstacles {
        k := sort.Search(len(lis), func(i int) bool { return lis[i] > o })
        if k == len(lis) {
            lis = append(lis, o)
        } else {
            lis[k] = o
        }
        res = append(res, k + 1)
    }
    return res
}