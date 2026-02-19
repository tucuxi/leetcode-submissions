func readBinaryWatch(turnedOn int) []string {
    var res []string
    var h, m uint
    for ; h < 12; h++ {
        hb := bits.OnesCount(h)
        for m = 0; m < 60; m++ {
            if hb + bits.OnesCount(m) == turnedOn {
                res = append(res, fmt.Sprintf("%d:%02d", h, m))
            }
        }
    }
    return res
}
