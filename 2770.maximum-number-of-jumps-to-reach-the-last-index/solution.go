func maximumJumps(nums []int, target int) int {
    c := make([]int, len(nums))
    c[0] = 1
    for i := range nums {
        if s := c[i]+1; s > 1 {    
            for j := i+1; j < len(nums); j++ {
                if d := nums[j]-nums[i]; -target <= d && d <= target && s > c[j] {
                    c[j] = s
                }
            }
        }
    }
    return c[len(c)-1]-1
}