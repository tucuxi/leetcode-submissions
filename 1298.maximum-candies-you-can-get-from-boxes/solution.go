func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
    res := 0
    boxes := make([]bool, len(status))
    var queue []int

    for _, b := range initialBoxes {
        if status[b] == 1 {
            queue = append(queue, b)
        } else {
            boxes[b] = true
        }
    }

    for len(queue) > 0 {
        i := queue[0]
        queue = queue[1:]
        res += candies[i]

        for _, b := range containedBoxes[i] {
            if status[b] == 0 {
                boxes[b] = true
            } else { 
                queue = append(queue, b)
            }
        }

        for _, b := range keys[i] {
            if status[b] == 0 {
                if boxes[b] {
                    boxes[b] = false
                    queue = append(queue, b)
                } else {
                    status[b] = 1
                }
            }
        }
    }

    return res
}