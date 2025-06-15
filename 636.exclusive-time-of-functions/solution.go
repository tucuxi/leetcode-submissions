func exclusiveTime(n int, logs []string) []int {
    res := make([]int, n)
    var stack []int
    var previousTimestamp int

    for _, log := range logs {
        s := strings.Split(log, ":")
        id, _ := strconv.Atoi(s[0])
        timestamp, _ := strconv.Atoi(s[2])
        top := len(stack) - 1

        switch s[1] {
            case "start":
                if top >= 0 {
                    res[stack[top]] += timestamp - previousTimestamp     
                }
                stack = append(stack, id)
                previousTimestamp = timestamp
            case "end":
                res[id] += timestamp + 1 - previousTimestamp
                stack = stack[:top]
                previousTimestamp = timestamp + 1
        }
    }
    return res
}