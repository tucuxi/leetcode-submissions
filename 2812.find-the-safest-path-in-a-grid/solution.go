func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return 0
    }

    md := make([][]int, n)
    var q [][2]int

    for i := range n {
        md[i] = make([]int, n)
        for j := range n {
            if grid[i][j] == 1 {
                q = append(q, [2]int{i, j})
                md[i][j] = 0
            } else {
                md[i][j] = -1
            }
        }
    }

    steps := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

    head := 0
    for head < len(q) {
        curr := q[head]
        head++
        i, j := curr[0], curr[1]

        for _, s := range steps {
            i2, j2 := i+s[0], j+s[1]
            if i2 >= 0 && i2 < n && j2 >= 0 && j2 < n && md[i2][j2] == -1 {
                md[i2][j2] = md[i][j] + 1
                q = append(q, [2]int{i2, j2})
            }
        }
    }

    var dfs func(int, int, int, [][]bool) bool
    dfs = func(i, j, sf int, v [][]bool) bool {
        if i2, j2 := i, j; md[i2][j2] < sf || v[i2][j2] {
            return false
        }
        if i == n-1 && j == n-1 {
            return true
        }
        v[i][j] = true
        for _, s := range steps {
            i2, j2 := i+s[0], j+s[1]
            if i2 >= 0 && i2 < n && j2 >= 0 && j2 < n { 
                if dfs(i2, j2, sf, v) {
                    return true
                }
            }
        }
        return false
    }

    low, high := 0, md[0][0]
    ans := 0

    for low <= high {
        mid := low + (high-low)/2
        
        v := make([][]bool, n)
        for i := range n {
            v[i] = make([]bool, n)
        }

        if dfs(0, 0, mid, v) {
            ans = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return ans
}