func haveConflict(event1 []string, event2 []string) bool {
    s1 := minutes(event1[0])
    e1 := minutes(event1[1])
    s2 := minutes(event2[0])
    e2 := minutes(event2[1])
    if s1 < s2 {
        return e1 >= s2
    }
    return e2 >= s1
}

func minutes(s string) int {
    h, _ := strconv.Atoi(s[:2])
    m, _ := strconv.Atoi(s[3:])
    return h * 60 + m
}