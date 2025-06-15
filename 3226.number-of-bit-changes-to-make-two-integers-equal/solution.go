const upper = 1e6

func minChanges(n int, k int) int {
    changes := 0
    for m := 1; m <= upper; m <<= 1 {
        nm := n & m != 0
        km := k & m != 0
        if nm == km {
            continue
        }
        if !nm && km {
            return -1
        }
        changes++
    }
    return changes
}