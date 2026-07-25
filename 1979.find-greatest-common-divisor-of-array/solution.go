func findGCD(nums []int) int {
    smallest, largest := optima(nums)
    return gcd(smallest, largest)
}

func optima(nums []int) (int, int) {
    smallest, largest := nums[0], nums[0]
    for _, num := range nums {
        smallest = min(num, smallest)
        largest = max(num, largest)
    }
    return smallest, largest   
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}