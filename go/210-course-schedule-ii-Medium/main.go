func findOrder(numCourses int, prerequisites [][]int) []int {
    in := make([]map[int]bool, numCourses)
    out := make([]map[int]bool, numCourses)
    for _, p := range prerequisites {
        if in[p[0]] == nil {
            in[p[0]] = make(map[int]bool)
        }
        in[p[0]][p[1]] = true
        if out[p[1]] == nil {
            out[p[1]] = make(map[int]bool)
        }
        out[p[1]][p[0]] = true
    }
    q := make([]int, 0, numCourses)
    for i := range in {
        if len(in[i]) == 0 {
            q = append(q, i)
        }
    }
    res := make([]int, 0, numCourses)
    for len(q) > 0 {
        res = append(res, q[0])
        for k := range out[q[0]] {
            delete(in[k], q[0])
            if len(in[k]) == 0 {
                q = append(q, k)
            }
        }
        q = q[1:]
    }
    if len(res) < numCourses {
        return []int{}
    }
    return res
}