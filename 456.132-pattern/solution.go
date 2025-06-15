func find132pattern(nums []int) bool {
    nj1 := math.MaxInt
    nj2 := math.MinInt
    ni := nums[0]
    for _, nk := range nums[1:] {
        if nk >= nj1 && nk <= nj2 {
            return true
        }
        nip := ni + 1
        nkm := nk - 1
        if nip <= nkm {
            if nip < nj1 {
                nj1 = nip
            }
            if nkm > nj2 {
                nj2 = nkm
            } 
        }
        if nk < ni {
            ni = nk
        }
    }
    return false
}