func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
	last := make([]int, m)
	j := m-1
	for i := n-1; i >= 0; i-- {
		if j >= 0 && word1[i] == word2[j] {
			last[j] = i
			j--
		}
	}

	res := make([]int, 0, m)
	skip := false
	j = 0

	for i := range n {
		if j == m {
			break
		}
		if word1[i] == word2[j] || (!skip && (j == m-1 || i < last[j+1])) {
			if word1[i] != word2[j] {
				skip = true
			}
			res = append(res, i)
			j++
		}
	}
	if j < m {
		return nil
	}
	return res
}