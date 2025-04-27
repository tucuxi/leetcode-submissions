func reverseBits(num uint32) uint32 {
    var res, mask uint32 = 0, 1
    for ; mask != 0; mask <<= 1 {
        res <<= 1
        if num & mask != 0 {
            res |= 1
        }
    }
    return res
}