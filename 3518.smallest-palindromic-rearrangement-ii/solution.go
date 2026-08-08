func smallestPalindrome(s string, k int) string {
	partition := len(s) / 2
	bucket := make([]int, 26)

	for i := range partition {
		bucket[s[i]-'a'] += 1
	}

	comb := func(n, m, kVal int) int {
		res := 1
        m = min(m, n-m)

		for i := 1; i <= m; i++ {
			res = res * (n - i + 1) / i
			if res > kVal {
				return kVal + 1
			}
		}
		return res
	}

	permutations := func(rem int) int {
		ways := 1
		for i := range bucket {
			if bucket[i] == 0 {
				continue
			}

			ways *= comb(rem, bucket[i], k)
			if ways > k {
				break
			}
			rem -= bucket[i]
		}
		return ways
	}

	left := []byte{}
	startIndex := 1

	for pos := 0; pos < partition; pos++ {
		for i := range bucket {
			if bucket[i] == 0 {
				continue
			}

			bucket[i]--

			ways := permutations(partition - pos - 1)
			if startIndex+ways > k {
				left = append(left, byte(i+'a'))
				break
			}

			bucket[i]++
			startIndex += ways
		}
	}

	if len(left) < partition {
		return ""
	}

	totalLen := len(s)
	res := make([]byte, totalLen)

	for i := range partition {
		res[i] = left[i]
		res[totalLen-1-i] = left[i]
	}

	if totalLen%2 != 0 {
		res[partition] = s[partition]
	}

	return string(res)
}