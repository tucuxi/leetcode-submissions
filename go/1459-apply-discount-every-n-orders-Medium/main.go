type Cashier struct {
    n int
    factor float64
    priceById map[int]int
    customers int   
}


func Constructor(n int, discount int, products []int, prices []int) Cashier {
    priceById := make(map[int]int)
    for i := range products {
        priceById[products[i]] = prices[i]
    }
    return Cashier{n, float64(100 - discount) / 100, priceById, 0}
}


func (this *Cashier) GetBill(product []int, amount []int) float64 {
    total := 0.
    for i := range product {
        total += float64(this.priceById[product[i]] * amount[i])
    }
    this.customers++
    if this.customers == this.n {
        this.customers = 0
        total *= this.factor
    }
    return total
}


/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */