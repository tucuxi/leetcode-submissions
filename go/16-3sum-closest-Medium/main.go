func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    diff := target - nums[0] - nums[1] - nums[2]
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            a := target - nums[i] - nums[j]
            off := j + 1
            k := off + sort.Search(len(nums) - off, func(i int) bool {
                return nums[off + i] >= a 
            })
            if k > off && abs(a - nums[k-1]) < abs(diff) {
                diff = a - nums[k-1]
            }
            if k < len(nums) && abs(a - nums[k]) < abs(diff) {
                diff = a - nums[k]
            }
        }
    }
    return target - diff
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}