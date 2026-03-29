func minAbsoluteDifference(nums []int) int {
	minDiff := -1
	last1, last2 := -1, -1

	for i, num := range nums {
		if num == 1 {
			last1 = i
			if last2 != -1 {
				diff := i - last2
				if minDiff == -1 || diff < minDiff {
					minDiff = diff
				}
			}
		} else if num == 2 {
			last2 = i
			if last1 != -1 {
				diff := i - last1
				if minDiff == -1 || diff < minDiff {
					minDiff = diff
				}
			}
		}
	}

	return minDiff    
}