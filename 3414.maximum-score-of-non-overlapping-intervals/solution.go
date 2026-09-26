func maximumWeight(intervals [][]int) []int {
	type Interval struct {
		l, r, weight, idx int
	}
 	n := len(intervals)
	arr := make([]Interval, n)
	for i := range arr {
		arr[i] = Interval{intervals[i][0], intervals[i][1], intervals[i][2], i}
	}
	// Sort by right endpoint.
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].r < arr[j].r
	})

	dp := make([][]int64, n+1)
	indices := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([]int64, 5)
		indices[i] = make([][]int, 5)
		for j := range 5 {
			indices[i][j] = []int{}
		}
	}

	for i := range arr {
		l, weight, idx := arr[i].l, arr[i].weight, arr[i].idx
		k := sort.Search(i, func(pos int) bool {
			return arr[pos].r >= l
		})

		for j := 1; j < 5; j++ {
			s1 := dp[i][j]
			s2 := dp[k][j-1] + int64(weight)
			if s1 > s2 {
				dp[i+1][j] = dp[i][j]
				indices[i+1][j] = append([]int{}, indices[i][j]...)
				continue
			}

			newIndex := append([]int{}, indices[k][j-1]...)
			newIndex = append(newIndex, idx)
			sort.Ints(newIndex)
			if s1 == s2 && compareSlices(indices[i][j], newIndex) < 0 {
				newIndex = append([]int{}, indices[i][j]...)
			}
			dp[i+1][j] = s2
			indices[i+1][j] = newIndex
		}
	}

	return indices[n][4]
}

func compareSlices(a, b []int) int {
	minLen := min(len(a), len(b))
	for i := range minLen {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return len(a) - len(b)
}