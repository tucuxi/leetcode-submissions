func minJumps(arr []int) int {
    m := map[int][]int{}
    for i := range arr {
        m[arr[i]] = append(m[arr[i]], i)
    }
    v := make([]bool, len(arr))
    v[0] = true
    j := 0
    for q := []int{0}; len(q) > 0; {
        for k := len(q); k > 0; k-- {
            i := q[0]
            q = q[1:]
            if i == len(arr) - 1 {
                return j
            }
            for _, next := range m[arr[i]] {
                if !v[next] {
                    v[next] = true
                    q = append(q, next)
                }
            }
            delete(m, arr[i])
            if next := i - 1; next >= 0 && !v[next] {
                v[next] = true
                q = append(q, next)
            }
            if next := i + 1; next < len(arr) && !v[next] {
                v[next] = true
                q = append(q, next)
            }
        }
        j++
    }
    panic("no solution")
}