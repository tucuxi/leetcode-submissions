type ProductOfNumbers struct {
    product []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{}
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.product = nil
    } else if len(this.product) == 0 {
        this.product = []int{num}
    } else {
        this.product = append(this.product, num * this.product[len(this.product) - 1])
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    if k > len(this.product) {
        return 0
    }
    if k == len(this.product) {
        return this.product[len(this.product) - 1]
    }
    return this.product[len(this.product) - 1] / this.product[len(this.product) - 1 - k]
}