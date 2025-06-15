func minimumIndex(nums []int) int {
    h := make(map[int]int)
    for _, num := range nums {
        h[num]++
    }

    dominant := 0
    for num, count := range h {
        if count > len(nums) / 2 {
            dominant = num
            break
        }
    }

    if dominant == 0 {
        return -1
    }

    l, r := 0, h[dominant]

    for i, num := range nums {
        if num == dominant {
            l++
            r--
        }
        if l > (i + 1) / 2 && r > (len(nums) - i - 1) / 2 {
            return i
        }
    }

    return -1
}