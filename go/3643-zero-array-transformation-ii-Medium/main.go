func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    sum := 0
    k := 0
    differences := make([]int, n + 1)

    for i, num := range nums {
        for sum + differences[i] < num {
            k++
            if k > len(queries) {
                return -1
            }
            l, r, v := queries[k - 1][0], queries[k - 1][1], queries[k - 1][2]
            if r >= i {
                differences[max(l, i)] += v
                differences[r + 1] -= v
            }
        }
        sum += differences[i]
    }
    return k
}