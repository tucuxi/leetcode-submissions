func lexGreaterPermutation(s string, target string) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	var res []byte
	n := len(target)

	for i := 0; i < n; i++ {
		targetChar := int(target[i] - 'a')

		// Case 1: First try to place the same character as target[i] at the
		// current position
		if cnt[targetChar] > 0 {
			cnt[targetChar]--
			// Check if the remaining characters can form a string greater than
			// target[i+1:]
			if canFormGreater(cnt, target, i+1) {
				res = append(res, target[i])
				continue
			}
			// Cannot form a larger string, backtrack
			cnt[targetChar]++
		}

		// Case 2: Place a character greater than target[i] at the current
		// position
		for j := targetChar + 1; j < 26; j++ {
			if cnt[j] > 0 {
				cnt[j]--
				res = append(res, byte('a'+j))
				// Fill remaining positions with the smallest lexicographical
				// order
				res = append(res, getMinString(cnt)...)
				return string(res)
			}
		}

		// No feasible solution found, return directly
		return ""
	}

	return ""
}

// Check if the remaining characters can form a string greater than the suffix.
func canFormGreater(cnt []int, target string, start int) bool {
	maxStr := getMaxString(cnt)
	suffix := target[start:]
	return maxStr > suffix
}

// Get the maximum lexicographical string (in descending order)
func getMaxString(cnt []int) string {
	var res []byte
	for i := 25; i >= 0; i-- {
		if cnt[i] > 0 {
			for k := 0; k < cnt[i]; k++ {
				res = append(res, byte('a'+i))
			}
		}
	}
	return string(res)
}

// Get the lexicographically smallest string (in ascending order)
func getMinString(cnt []int) string {
	var res []byte
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			for k := 0; k < cnt[i]; k++ {
				res = append(res, byte('a'+i))
			}
		}
	}
	return string(res)
}