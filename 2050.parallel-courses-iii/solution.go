func minimumTime(n int, relations [][]int, time []int) int {
    in := make([]map[int]bool, n)
    out := make([]map[int]bool, n)
    for _, r := range relations {
        prevCourse, nextCourse := r[0] - 1, r[1] - 1
        if in[nextCourse] == nil {
            in[nextCourse] = make(map[int]bool)
        }
        in[nextCourse][prevCourse] = true
        if out[prevCourse] == nil {
            out[prevCourse] = make(map[int]bool)
        }
        out[prevCourse][nextCourse] = true
    }
    nextCourses := make([]int, 0)
    completionTime := make([]int, n)
    for course, prereqs := range in {
        if len(prereqs) == 0 {
            nextCourses = append(nextCourses, course)
            completionTime[course] = time[course]
        }
    }
    for len(nextCourses) > 0 {
        course := nextCourses[0]
        nextCourses = nextCourses[1:]
        for next := range out[course] {
            completionTime[next] = max(completionTime[next], time[next] + completionTime[course])
            delete(in[next], course)
            if len(in[next]) == 0 {
                nextCourses = append(nextCourses, next)
            }
        }
    }
    return max(completionTime...)
}

func max(as ...int) int {
    res := as[0]
    for _, a := range as {
        if a > res {
            res = a
        }
    }
    return res
}