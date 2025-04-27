func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    zz := make([][]byte, numRows)
    r, c := 0, 0
    dr, dc := 1, 0
    for i := range s {
        zz[r] = append(zz[r], s[i])
        if r + dr == numRows {
            dr, dc = -1, 1
        } else if r + dr < 0 {
            dr, dc = 1, 0
        }
        r += dr
        c += dc
    }
    var res bytes.Buffer
    for i := range zz {
        res.Write(zz[i])
    }
    return res.String()
}