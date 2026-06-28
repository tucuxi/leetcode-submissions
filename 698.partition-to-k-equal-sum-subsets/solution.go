func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0
    
    for _, num := range nums {
        sum += num
    }

    if sum % k != 0 {
        return false
    }

    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    target := sum/k
    if nums[0] > target {
        return false
    }

    return search(nums, 0, make([]int, k), sum / k)
}

func search(nums []int, index int, sums []int, target int) bool {
    if index == len(nums) {
        return true
    }
    for i := range sums {
        if i > 0 && sums[i] == sums[i-1] {
            continue
        }
        sums[i] += nums[index]
        if sums[i] <= target {
            if search(nums, index+1, sums, target) {
                return true
            }
        }
        sums[i] -= nums[index]
    }
    return false
}