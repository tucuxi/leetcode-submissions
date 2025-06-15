const initial = "0000"

func openLock(deadends []string, target string) int {
    visited := make(map[string]bool)
    for _, d := range deadends {
        visited[d] = true
    }
    if visited[initial] {
        return -1
    }
    visited[initial] = true
    queue := []string{initial}
    res := 0
    for len(queue) > 0 {
        for i := len(queue); i > 0; i-- {
            state := queue[0]
            queue = queue[1:]
            if state == target {
                return res
            }
            for j := range state {
                nextUp := []byte(state)
                if nextUp[j] == '9' {
                    nextUp[j] = '0'
                } else {
                    nextUp[j]++
                }
                if s := string(nextUp); !visited[s] {
                    visited[s] = true
                    queue = append(queue, s)
                }
                nextDown := []byte(state)
                if nextDown[j] == '0' {
                    nextDown[j] = '9'
                } else {
                    nextDown[j]--
                }
                if s := string(nextDown); !visited[s] {
                    visited[s] = true
                    queue = append(queue, s)
                }
            }
        }
        res++
    }
    return -1
}