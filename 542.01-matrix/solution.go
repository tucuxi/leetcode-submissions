func updateMatrix(mat [][]int) [][]int {
    answer := make([][]int, len(mat))
    queue := [][2]int{}
    for i := range mat {
        answer[i] = make([]int, len(mat[i]))
        for j := range mat[i] {
            if mat[i][j] == 0 {
                queue = append(queue, [2]int{i, j})
            } else {
                answer[i][j] = math.MaxInt
            }
        }
    }
    for ; len(queue) > 0; queue = queue[1:] {
        r, c := queue[0][0], queue[0][1]
        for _, nb := range [][2]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
            if nb[0] >= 0 && nb[0] < len(mat) && nb[1] >= 0 && nb[1] < len(mat[nb[0]]) {
                if answer[r][c] + 1 < answer[nb[0]][nb[1]] {
                    answer[nb[0]][nb[1]] = answer[r][c] + 1
                    queue = append(queue, nb)
                }
            }
        }
    }
    return answer
}