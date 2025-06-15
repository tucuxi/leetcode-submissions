func minCost(arr []int, brr []int, k int64) int64 {
    opt1 := int64(0)
    for i := range arr {
        opt1 += int64(x(arr[i], brr[i]))
    }
    
	sort.Ints(arr)
    sort.Ints(brr)
    
    opt2 := k
    for i := range arr {
        opt2 += int64(x(arr[i], brr[i]))
    }
    
    return min(opt1, opt2)
}

func x(a, b int) int {
    if b - a >= 0 {
        return b - a
    }
    return a - b
}