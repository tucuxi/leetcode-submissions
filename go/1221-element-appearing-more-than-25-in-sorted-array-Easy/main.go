func findSpecialInteger(arr []int) int {
    sl := len(arr) / 4
    p := arr[0]
    for i := 1; i < 4; i++ {
        q := arr[i * sl]
        if q == p {
            return p
        }
        p = q
    }
    for i := 0; i < len(arr); i += sl {
        if cont(arr, arr[i]) > sl {
            return arr[i]
        } 
    }
    panic("not found")
}

func cont(arr []int, n int) int {
    a, b := 0, len(arr) - 1
    for a < b {
        m := (a + b) / 2
        if arr[m] < n {
            a = m + 1
        } else {
            b = m
        }
    }
    first := a
    b = len(arr) - 1
    for a < b {
        m := (a + b) / 2 + 1
        if arr[m] > n {
            b = m - 1
        } else {
            a = m
        }
    }
    return b - first + 1
}