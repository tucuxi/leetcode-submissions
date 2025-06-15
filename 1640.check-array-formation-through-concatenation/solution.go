func canFormArray(arr []int, pieces [][]int) bool {
    piecesMap := map[int]int{}
    for i := range pieces {
        piecesMap[pieces[i][0]] = i
    }
    p := []int{}
    for i := range arr {
        if len(p) == 0 {
            if _, ok := piecesMap[arr[i]]; !ok {
                return false
            }
            p = pieces[piecesMap[arr[i]]]
        }
        if arr[i] != p[0] {
            return false
        }
        p = p[1:]
    }
    return true
}