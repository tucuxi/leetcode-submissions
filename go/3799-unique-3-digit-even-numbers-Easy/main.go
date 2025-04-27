func totalNumbers(digits []int) int {
    nums := make(map[int]bool)
    for i := range digits {
        if digits[i] == 0 {
            continue
        }
        for j := range digits {
            if j == i {
                continue
            }
            for k := range digits {
                if k == i || k == j || digits[k] % 2 != 0 {
                    continue
                }
                nums[100 * digits[i] + 10 * digits[j] + digits[k]] = true
            }
        }
    }
    return len(nums)
}