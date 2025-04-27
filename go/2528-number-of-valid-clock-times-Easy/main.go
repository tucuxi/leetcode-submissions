func countTime(time string) int {
    res := 1
    if time[:2] == "??" {
        res *= 24
    } else if time[1:2] == "?" {
        if time[:1] == "2" {
            res *= 4
        } else {
            res *= 10
        }
    } else if time[:1] == "?" {
        if time[1:2] < "4" {
            res *= 3
        } else {
            res *= 2
        }
    }
    if time[3:4] == "?" {
        res *= 6
    }
    if time [4:] == "?" {
        res *= 10
    }
    return res
}