func majorityElement(nums []int) int {
    var major, count int 
    
    for _, n := range nums {
        if count == 0 {
            major = n
            count = 1
        } else if n == major {
            count++
        } else {
            count--
        }
    }
    return major
}