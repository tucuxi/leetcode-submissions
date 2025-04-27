const cores = 4

func minProcessingTime(processorTime []int, tasks []int) int {
    sort.Ints(processorTime)
    sort.Ints(tasks)
    
    l := len(tasks) - 1
    t := 0
    
    for i := range tasks {
        c := processorTime[i/cores] + tasks[l-i]
        if c > t {
            t = c
        }
    }
    return t
}