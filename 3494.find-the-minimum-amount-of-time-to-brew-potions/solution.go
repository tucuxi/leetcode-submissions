func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	times := make([]int64, n)

	for j := range m {
		var curTime int64 = 0
		for i := range n {
			if curTime < times[i] {
				curTime = times[i]
			}
			curTime += int64(skill[i]) * int64(mana[j])
		}
		times[n-1] = curTime
		for i := n - 2; i >= 0; i-- {
			times[i] = times[i+1] - int64(skill[i+1])*int64(mana[j])
		}
	}
	return times[n-1]
}