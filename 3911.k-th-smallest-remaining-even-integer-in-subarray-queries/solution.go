func kthRemainingInteger(nums []int, queries [][]int) []int {
    var evenIndices []int

    for i, num := range nums {
        if num%2 == 0 {
            evenIndices = append(evenIndices, i)
        }
    }

    ans := make([]int, len(queries))

    for i, q := range queries {
        l, r, k := q[0], q[1], q[2]

        start := sort.Search(len(evenIndices), func(j int) bool { return evenIndices[j] >= l })
        end := sort.Search(len(evenIndices), func(j int) bool { return evenIndices[j] > r })

        lo, hi := 1, end-start
        best := 0

        for lo <= hi {
            mid := lo + (hi-lo)/2
            num := nums[evenIndices[start+mid-1]]
            if num <= 2*(k+mid-1) {
                best = mid
                lo = mid+1
            } else {
                hi = mid-1
            }
        }
        ans[i] = 2*(k+best)
    }

    return ans
}