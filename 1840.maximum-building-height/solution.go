func maxBuilding(n int, restrictions [][]int) int {
	if len(restrictions) == 0 {
		return n - 1
	}

	slices.SortFunc(restrictions, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	m := len(restrictions)

	if restrictions[0][1] > restrictions[0][0]-1 {
		restrictions[0][1] = restrictions[0][0] - 1
	}

	for i := 1; i < m; i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]
		if restrictions[i][1] > restrictions[i-1][1]+dist {
			restrictions[i][1] = restrictions[i-1][1] + dist
		}
	}

	for i := m - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]
		if restrictions[i][1] > restrictions[i+1][1]+dist {
			restrictions[i][1] = restrictions[i+1][1] + dist
		}
	}

	maxHeight := 0
	prevId, prevH := 1, 0

	for i := 0; i < m; i++ {
		currId, currH := restrictions[i][0], restrictions[i][1]
		peak := (prevH + currH + (currId - prevId)) / 2

		if peak > maxHeight {
			maxHeight = peak
		}

		prevId, prevH = currId, currH
	}

	finalPeak := prevH + (n - prevId)
	if finalPeak > maxHeight {
		maxHeight = finalPeak
	}

	return maxHeight
}