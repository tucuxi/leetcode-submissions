func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)

	countSubarrays := func(target int) int64 {
		var count int64 = 0
		l := 0
		maxQ := make([]int, 0)
		minQ := make([]int, 0)

		for r := 0; r < n; r++ {
			for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[r] {
				maxQ = maxQ[:len(maxQ)-1]
			}
			maxQ = append(maxQ, r)

			for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[r] {
				minQ = minQ[:len(minQ)-1]
			}
			minQ = append(minQ, r)

			for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] >= target {
				count += int64(n - r)
				if maxQ[0] == l {
					maxQ = maxQ[1:]
				}
				if minQ[0] == l {
					minQ = minQ[1:]
				}
				l++
			}
		}
		return count
	}

	var low, high, bestThreshold int = 0, 1e9, 0
	for low <= high {
		mid := low + (high-low)/2
		if countSubarrays(mid) >= int64(k) {
			bestThreshold = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	var totalVal int64 = 0
	var exactCount int64 = 0
	l := 0
	target := bestThreshold + 1
	maxQ := make([]int, 0)
	minQ := make([]int, 0)

	for r := 0; r < n; r++ {
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[r] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)

		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[r] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)

		for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] >= target {
			curMax := nums[maxQ[0]]
			curMin := nums[minQ[0]]

			for rPrime := r; rPrime < n; rPrime++ {
				if nums[rPrime] > curMax {
					curMax = nums[rPrime]
				}
				if nums[rPrime] < curMin {
					curMin = nums[rPrime]
				}
				totalVal += int64(curMax - curMin)
				exactCount++
			}

			if maxQ[0] == l {
				maxQ = maxQ[1:]
			}
			if minQ[0] == l {
				minQ = minQ[1:]
			}
			l++
		}
	}

	return totalVal + (int64(k)-exactCount)*int64(bestThreshold)
}