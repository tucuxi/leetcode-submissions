type arrayValue struct {
    value int
    index int
}

func maxDistance(arrays [][]int) int {
    minimum1, minimum2 := arrayValue{math.MaxInt, -1}, arrayValue{math.MaxInt, -1}
    maximum1, maximum2 := arrayValue{math.MinInt, -1}, arrayValue{math.MinInt, -1}
    for index, array := range arrays {
        if v := array[0]; v < minimum1.value {
            minimum2 = minimum1
            minimum1.value = v
            minimum1.index = index
        } else if v < minimum2.value {
            minimum2.value = v
            minimum2.index = index
        }
        if v := array[len(array) - 1]; v > maximum1.value {
            maximum2 = maximum1
            maximum1.value = v
            maximum1.index = index
        } else if v > maximum2.value {
            maximum2.value = v
            maximum2.index = index
        }
    }
    if minimum1.index == maximum1.index {
        return max(maximum2.value - minimum1.value, maximum1.value - minimum2.value)
    }
    return maximum1.value - minimum1.value
}