func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    m := len(students)
    n := len(students[0])
    c := make([][]int, m)

    for i := range m {
        c[i] = make([]int, m)
        for j := range m {
            for k := range n {
                if students[i][k] == mentors[j][k] {
                    c[i][j]++
                }
            }
        }
    }

    res := 0
    v := make([]bool, m) 

    var p func(studentIdx int, currentSum int)
    
    p = func(studentIdx int, currentSum int) {
        if studentIdx == m {
            res = max(currentSum, res)
            return
        }
        for j := range m {
            if !v[j] {
                v[j] = true
                p(studentIdx+1, currentSum + c[studentIdx][j])
                v[j] = false
            }
        }
    }

    p(0, 0)
    return res
}