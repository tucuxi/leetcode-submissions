func restoreArray(adjacentPairs [][]int) []int {
    neighbors := func() map[int][]int {
        res := make(map[int][]int)
        for _, p := range adjacentPairs {
            res[p[0]] = append(res[p[0]], p[1])
            res[p[1]] = append(res[p[1]], p[0])
        }
        return res
    }()
    nums := make([]int, len(adjacentPairs) + 1)
    nums[0], nums[len(nums) - 1] = func() (int, int) {
        res := make([]int, 0, 2)
        for n, nb := range neighbors {
            if len(nb) == 1 {
                res = append(res, n)
            }
        }
        return res[0], res[1]
    }()
    for i := 1; i < len(nums) - 1; i++ {
        k := neighbors[nums[i - 1]][0]
        if neighbors[k][0] == nums[i - 1] {
            neighbors[k] = neighbors[k][1:]
        } else {
            neighbors[k] = neighbors[k][:1]
        }
        nums[i] = k
    }
    return nums
}