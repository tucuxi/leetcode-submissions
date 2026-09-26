func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    set := make(map[[2]int]struct{})
    list := [][2]int{}
    for i := range n {
        for j := range n {
            if img2[i][j] == 1 {
                set[[2]int{i, j}] = struct{}{}
            }
            if img1[i][j] == 1 {
                list = append(list, [2]int{i, j})
            }
        }
    }
    res := 0
    for i := -n+1; i <= n-1; i++ {
        for j := -n+1; j <= n-1; j++ {
            overlap := 0
            for _, v := range list {
                ii := v[0] + i
                jj := v[1] + j
                if ii < 0 || jj < 0 || ii >= n || jj >= n {
                    continue
                }
                if _, ok := set[[2]int{ii, jj}]; ok {
                    overlap++
                }
            }
            if overlap > res {
                res = overlap
            }
        }
    }
    return res
}