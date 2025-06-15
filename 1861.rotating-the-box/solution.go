func rotateTheBox(box [][]byte) [][]byte {
    m, n := len(box), len(box[0])
    
    r := make([][]byte, n)
    for i := range n {
        r[i] = make([]byte, m)
        for j := range m {
            r[i][j] = box[m - 1 - j][i]
        }
    }

    for j := range m {
        for i := n - 2; i >= 0; i-- {
            for k := i + 1; k < n; k++ {
                if r[k - 1][j] == '#' && r[k][j] == '.' {
                    r[k - 1][j] = '.'
                    r[k][j] = '#'
                }
            }
        }
    }

    return r
}

/*
1 2 3 
4 5 6
7 8 9

7 4 1
8 5 2 
9 6 3

(0, 0) -> (0, 2)
(0, 1) -> (1, 2)
(0, 2) -> (2, 2)
(1, 0) -> (0, 1)
(1, 1) -> (1, 1)
(1, 2) -> (2, 1)
(2, 0) -> (0, 0)
(2, 1) -> (1, 0)
(2, 2) -> (2, 0)

(i, j) <- (m - 1 - j, i)
*/