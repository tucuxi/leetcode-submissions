func findKthNumber(n int, k int) int {
    curr := 1
	for j := k - 1; j > 0; {
		step := calStep(n, curr, curr + 1)
		if step <= j {
			curr++
            j -= step
		} else {
			curr *= 10
            j--
		}
	}
	return curr
}

func calStep(n, prefix1, prefix2 int) int {
	res := 0
	for prefix1 <= n {
		res += min(n + 1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}
	return res
}