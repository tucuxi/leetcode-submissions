func majorityElement(nums []int) []int {
    candidate1 := 0
    candidate2 := 1
    count1 := 0
    count2 := 0
    for _, n := range nums {
        if n == candidate1 {
            count1++
        } else if n == candidate2 {
            count2++
        } else if count1 == 0 {
            candidate1 = n
            count1 = 1
        } else if count2 == 0 {
            candidate2 = n
            count2 = 1
        } else {
            count1--
            count2--
        }
    }
    
    var res []int
    for _, candidate := range []int{candidate1, candidate2} {
        count := 0
        for _, n := range nums {
            if n == candidate {
                count++
            }
        }
        if count > len(nums) / 3 {
            res = append(res, candidate)
        }
    }
    return res
}