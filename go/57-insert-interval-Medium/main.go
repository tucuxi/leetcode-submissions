// newInterval:               |-----------|
// a. intervals:  |--|  |-||
// b. intervals:     |-----|                   |****|  |----|  ...
// c. intervals: |--|   |**********|  |+++++|  |----|  |----|  ...
// d. intervals:     |---|     |***|  |+|      |----|  |----|   ...

func insert(intervals [][]int, newInterval []int) [][]int {
    k := sort.Search(len(intervals), func(i int) bool {
        return intervals[i][1] >= newInterval[0]
    })
    if k == len(intervals) {    // case a
        return append(intervals, newInterval)
    }
    if intervals[k][0] > newInterval[1] {    // case b
        intervals = append(intervals, intervals[len(intervals) - 1])
        copy(intervals[k+1:], intervals[k:])
        intervals[k] = newInterval
        return intervals
    }
    if intervals[k][0] > newInterval[0] {    // case d
        intervals[k][0] = newInterval[0]
    }
    if intervals[k][1] < newInterval[1] {    // case c
        intervals[k][1] = newInterval[1]
    }
    i := k + 1
    for i < len(intervals) && intervals[i][0] <= intervals[k][1] {
        if intervals[i][1] > intervals[k][1] {
            intervals[k][1] = intervals[i][1]
        }
        i++
    }
    return append(intervals[:k + 1], intervals[i:]...)
}