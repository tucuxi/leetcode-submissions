func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	pre := make([]int, 2*n+1)
	pre[n] = 1
	cnt := n
	ans := int64(0)
    presum := int64(0)
	for i := range n {
		if nums[i] == target {
			presum += int64(pre[cnt])
			cnt++
			pre[cnt]++
		} else {
			cnt--
			presum -= int64(pre[cnt])
			pre[cnt]++
		}
		ans += presum
	}
	return ans
}