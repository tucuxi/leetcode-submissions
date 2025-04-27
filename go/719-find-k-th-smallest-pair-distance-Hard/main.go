func smallestDistancePair(nums []int, k int) int {

    n := len(nums)

    countPairs := func(dist int) int {
        count := 0
        left := 0
        for right := 1; right < n; right++ {
            for left < right && nums[right] - nums[left] > dist {
                left++
            }
            count += right - left
        }
        return count
    }

    slices.Sort(nums)

    return sort.Search(nums[n-1] - nums[0] + 1, func(dist int) bool {
        return countPairs(dist) >= k
    })
}