func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    g := make([][]bool, numCourses)
    for i := range g {
        g[i] = make([]bool, numCourses)
    }
    for _, p := range prerequisites {
        g[p[0]][p[1]] = true
    }
    for i := range numCourses {
        for j := range numCourses {
            if g[j][i] {
                for k := range numCourses {
                    g[j][k] = g[j][k] || g[i][k]
                } 
            }
        }
    }
    res := make([]bool, 0, len(queries))
    for _, q := range queries {
        res = append(res, g[q[0]][q[1]])
    }
    return res
}