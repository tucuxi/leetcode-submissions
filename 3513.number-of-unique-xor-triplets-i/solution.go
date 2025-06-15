func uniqueXorTriplets(nums []int) int {
    switch n := len(nums); n {
        case 1, 2:
            return n
        default:
            return 1 << (bits.UintSize - bits.LeadingZeros(uint(n)))
    }
}
