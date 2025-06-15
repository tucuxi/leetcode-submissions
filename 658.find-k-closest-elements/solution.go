func findClosestElements(arr []int, k int, x int) []int {
    p, q := 0, len(arr) - k
    for p < q {
        m := p + (q - p) / 2
        if x - arr[m] > arr[m + k] - x {
            p = m + 1
        } else {
            q = m
        }
    }
    return arr[p:p+k]
  }