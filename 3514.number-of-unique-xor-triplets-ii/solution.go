func uniqueXorTriplets(nums []int) int {
    var pairs, triplets [2048]bool

    for i, a := range nums {
        for _, b := range nums[i:] {
            pairs[a^b] = true
        }
    }

    for i, p := range pairs {
        if p {
            for _, num := range nums {
                triplets[i^num] = true
            }
        }
    }

    res := 0

    for _, b := range triplets {
        if b {
            res++
        }
    }
    
    return res
}