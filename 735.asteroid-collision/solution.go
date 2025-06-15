func asteroidCollision(asteroids []int) []int {
    st := make([]int, 0, len(asteroids))
    
    outer:
    for _, a := range asteroids {
        if a < 0 {
            for len(st) > 0 {
                b := st[len(st) - 1]
                if b < 0 {
                    break
                }
                if -a < b {
                    continue outer
                }
                st = st[:len(st) - 1]
                if -a == b {
                    continue outer
                }
            }
        }
        st = append(st, a)
    }
    return st
}