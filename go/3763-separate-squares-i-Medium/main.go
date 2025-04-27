func separateSquares(squares [][]int) float64 {
	totalArea := 0.0
	for _, square := range squares {
		li := float64(square[2])
		totalArea += li * li
	}

	left, right := 0.0, 1e14
	precision := 1e-5
	for right-left > precision {
		mid := (left + right) / 2
		areaBelow := 0.0
		for _, square := range squares {
			yi, li := float64(square[1]), float64(square[2])
			if yi+li <= mid {
				areaBelow += li * li
			} else if yi < mid {
				heightBelow := mid - yi
				if heightBelow > li {
					heightBelow = li
				}
				areaBelow += heightBelow * li
			}
		}

		if areaBelow*2 < totalArea {
			left = mid
		} else {
			right = mid
		}
	}

	return (left + right) / 2
}