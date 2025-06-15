func sortColors(nums []int)  {
    var count [3]int
    
    for _, n := range nums {
        count[n]++
    }
    nums = nums[:0]
    for i, n := range count {
        for range n {
            nums = append(nums, i)
        }
    }
}