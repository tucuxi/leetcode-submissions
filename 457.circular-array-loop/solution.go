func circularArrayLoop(nums []int) bool {
    
    advance := func(start, current, steps int) (int, bool) {
        p := current
        for range steps {
            if nums[start] * nums[p] < 0 {
                return -1, false
            }
            q := ((p + nums[p]) % len(nums) + len(nums)) % len(nums)
            if q == p {
                return -1, false
            }
            p = q
        }
        return p, true
    }

    for i := range nums {
        slow, fast := i, i
        for {
            slow2, slowOk := advance(i, slow, 1)
            fast2, fastOk := advance(i, fast, 2)
            if !slowOk || !fastOk {
                break
            }
            if slow2 == fast2 {
                return true
            }
            slow, fast = slow2, fast2
        }
    }
    return false
}