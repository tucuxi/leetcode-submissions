func fullBloomFlowers(flowers [][]int, people []int) []int {
    changes := func() []int {
        res := make([]int, 0, 2 * len(flowers))
        for _, f := range flowers {
            res = append(res, f[0], -(f[1] + 1))
        }
        sort.Slice(res, func(i, j int) bool {
            return abs(res[i]) < abs(res[j])
        })
        return res
    }()
    idx := func() []int {
        res := make([]int, len(people))
        for i := range res {
            res[i] = i
        }
        sort.Slice(res, func(i, j int) bool {
            return people[res[i]] < people[res[j]]
        })
        return res
    }()
    return func() []int {
        res := make([]int, len(people))
        cur := 0
        i, k := 0, 0
        for k < len(people) {
            p := idx[k]
            if i == len(changes) || people[p] < abs(changes[i]) {
                res[p] = cur
                k++
            } else {
                cur += sgn(changes[i])
                i++
            }
        }
        return res
    }()
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func sgn(a int) int {
    switch {
        case a > 0: return 1
        case a < 0: return -1
        default: return 0
    }
}