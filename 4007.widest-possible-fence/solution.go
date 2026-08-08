func maximumWidth(planks []int) int {
	freq := make(map[int]int)
	for _, p := range planks {
		freq[p]++
	}

	uniquePlanks := make([]int, 0, len(freq))
	for p := range freq {
		uniquePlanks = append(uniquePlanks, p)
	}

	widthForHeight := make(map[int]int)

	for _, p := range uniquePlanks {
		widthForHeight[p] += freq[p]
	}

	for i, p1 := range uniquePlanks {
		if freq[p1] >= 2 {
			widthForHeight[p1*2] += freq[p1] / 2
		}

		for j := i + 1; j < len(uniquePlanks); j++ {
			p2 := uniquePlanks[j]
			sum := p1 + p2
			pairs := min(freq[p1], freq[p2])
			widthForHeight[sum] += pairs
		}
	}

	maxWidth := 0
	for _, width := range widthForHeight {
		maxWidth = max(width, maxWidth)
	}

	return maxWidth
}