func rand10() int {
    for {
        r := 7 * (rand7() - 1) + rand7() - 1
        if r < 40 {
            return r % 10 + 1
        }
    }
}