func canFinish(numCourses int, prerequisites [][]int) bool {
    successors := make([][]int, numCourses)
    predecessors := make([]map[int]bool, numCourses)
    for _, req := range prerequisites {
        successors[req[1]] = append(successors[req[1]], req[0])
        if predecessors[req[0]] == nil {
            predecessors[req[0]] = map[int]bool{req[1]: true}
        } else {
            predecessors[req[0]][req[1]] = true
        }
    }
    queue := []int{}
    for course, pre := range predecessors {
        if len(pre) == 0 {
            queue = append(queue, course)
        }
    }
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        numCourses--
        for _, suc := range successors[course] {
            delete(predecessors[suc], course)
            if len(predecessors[suc]) == 0 {
                queue = append(queue, suc)
            }
        }
    }
    return numCourses == 0
}