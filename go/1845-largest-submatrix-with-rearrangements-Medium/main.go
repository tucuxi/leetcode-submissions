func largestSubmatrix(matrix [][]int) int {
    res := 0
    n := len(matrix[0])
    prevRow := make([]int, n)
    
    for i := range matrix {
        currRow := make([]int, n)
        copy(currRow, matrix[i])
        for j := range currRow {
            if currRow[j] != 0 {
                currRow[j] += prevRow[j]
            }
        }
        
        copy(prevRow, currRow)
        
        sort.Ints(currRow)
        for j := range currRow {
            if a := currRow[j] * (n - j); a > res {
                res = a
            }
        }
    }
    
    return res
}