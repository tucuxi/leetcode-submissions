func generateString(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	s := make([]byte, n+m-1)
	fixed := make([]int, n+m-1)

	for i := range s {
		s[i] = 'a'
	}

	for i := range n {
		if str1[i] == 'T' {
			for j := i; j < i+m; j++ {
				if fixed[j] == 1 && s[j] != str2[j-i] {
					return ""
				}
                s[j] = str2[j-i]
				fixed[j] = 1
			}
		}
	}

    outer:
	for i := range n {
		if str1[i] == 'F' {
			idx := -1
			for j := i + m - 1; j >= i; j-- {
				if str2[j-i] != s[j] {
					continue outer
				}
				if idx == -1 && fixed[j] == 0 {
					idx = j
				}
			}
            if idx == -1 {
                return ""
            }
			s[idx] = 'b'
		}
	}
	return string(s)
}