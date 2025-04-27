func uniqueXorTriplets(nums []int) int {
    if n := len(nums); n <= 2 {
        return n
    }
    
    var pairs [2048]bool
    for i, a := range nums {
        for _, b := range nums[i + 1:] {
            pairs[a ^ b] = true
        }
    }

    res := 0
    var triplets [2048]bool
    for i, p := range pairs {
        if p {
            for _, num := range nums {
                x := i ^ num
                if !triplets[x] {
                    triplets[x] = true
                    res++
                }
            }
        }
    }

    return res
}