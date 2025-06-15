func minDominoRotations(tops []int, bottoms []int) int {
    res := len(tops)
    for i := 1; i <= 6; i++ {
        if r, ok := check(i, tops, bottoms); ok {
            res = min(res, r)
        }
        if r, ok :=check(i, bottoms, tops); ok {
            res = min(res, r)
        }
    }
    if res == len(tops) {
        return -1
    }
    return res
}

func check(num int, row1, row2 []int) (int, bool) {
    res := 0
    for i := range row1 {
        if row1[i] == num {
            continue
        } else if row2[i] == num {
            res++
        } else {
            return -1, false
        }
    }
    return res, true
}