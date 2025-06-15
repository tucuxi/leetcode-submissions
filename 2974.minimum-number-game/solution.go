func numberGame(nums []int) []int {
    var c [101]int
    
    for _, n := range nums {
        c[n]++
    }
    
    arr := make([]int, len(nums))
    
    for i, k := 1, 0; k < len(arr); k++ {
        for ; c[i] == 0; i++ {}
        c[i]--
        arr[k^1] = i
    }
    return arr
}