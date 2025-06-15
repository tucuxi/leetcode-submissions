func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
    if a == e {
        if c != e || b < f && (d < b || d > f) || b > f && (d > b || d < f) {
            return 1
        }
    }
    if b == f {
        if d != f || a < e && (c < a || c > e) || a > e && (c > a || c < e) {
            return 1
        }
    }
    dc, dd := sgn(e-c), sgn(f-d)
    for {
        c += dc
        d += dd
        if c < 1 || c > 8 || d < 1 || d > 8 {
            break
        }
        if c == a && d == b {
            break
        }
        if c == e && d == f {
            return 1
        }
    }
    return 2
}

func sgn(a int) int {
    if a < 0 {
        return -1
    }
    return 1
}
/*
1 ..Q
2 .R.
3 B..
  123

a, b = 2, 2
c, d = 3, 1
e, f = 1, 3

dc, df = -1, 1

c, d = 2, 2
*/
