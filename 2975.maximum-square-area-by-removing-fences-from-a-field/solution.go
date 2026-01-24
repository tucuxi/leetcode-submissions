const MOD = 1000000007

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hEdges := getEdges(hFences, m)
	vEdges := getEdges(vFences, n)

	var res int64

	for e := range hEdges {
		if _, exists := vEdges[e]; exists {
			if int64(e) > res {
				res = int64(e)
			}
		}
	}

	if res == 0 {
		return -1
	}
	return int((res * res) % MOD)
}

func getEdges(fences []int, border int) map[int]bool {
	set := make(map[int]bool)
	list := make([]int, len(fences), len(fences) + 2)

	copy(list, fences)
	list = append(list, 1, border)
	slices.Sort(list)

	for i := range list {
		for j := i + 1; j < len(list); j++ {
			set[list[j]-list[i]] = true
		}
	}

	return set
}