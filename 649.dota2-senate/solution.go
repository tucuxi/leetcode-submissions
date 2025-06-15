func predictPartyVictory(senate string) string {
    dq, rq := []int{}, []int{}
    for i, s := range senate {
        if s == 'D' {
            dq = append(dq, i)
        } else {
            rq = append(rq, i)
        }
    }
    for len(dq) > 0 && len(rq) > 0 {
        if dq[0] < rq[0] {
            dq = append(dq[1:], dq[0] + len(senate))
            rq = rq[1:]
        } else {
            dq = dq[1:]
            rq = append(rq[1:], rq[0] + len(senate))
        }
    }
    if len(dq) < len(rq) {
        return "Radiant"
    }
    return "Dire"
}