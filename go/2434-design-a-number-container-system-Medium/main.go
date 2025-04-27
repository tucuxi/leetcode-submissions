type NumberContainers struct {
    indexToNumber map[int]int
    numberToIndex map[int][]int
    isSorted      map[int]bool
}


func Constructor() NumberContainers {
    return NumberContainers{
        indexToNumber: make(map[int]int),
        numberToIndex: make(map[int][]int),
        isSorted:      make(map[int]bool),
    }
}


func (this *NumberContainers) Change(index int, number int)  {
    this.indexToNumber[index] = number
    this.numberToIndex[number] = append(this.numberToIndex[number], index)
    this.isSorted[number] = false
}


func (this *NumberContainers) Find(number int) int {
    if !this.isSorted[number] {
        slices.Sort(this.numberToIndex[number])
        this.isSorted[number] = true
    }
    for _, index := range this.numberToIndex[number] {
        if this.indexToNumber[index] == number {
            return index
        }
    }

    return -1
}