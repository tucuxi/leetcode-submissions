type ATM struct {
    banknoteCounts []int
}

var denominations = []int{20, 50, 100, 200, 500}

func Constructor() ATM {
    return ATM{make([]int, len(denominations))}
}


func (this *ATM) Deposit(banknotesCount []int)  {
    for i, b := range banknotesCount {
        this.banknoteCounts[i] += b
    }
}


func (this *ATM) Withdraw(amount int) []int {
    b := make([]int, len(denominations))
    for i := len(b) - 1; i >= 0; i-- {
        b[i] = min(this.banknoteCounts[i], amount / denominations[i])
        amount -= b[i] * denominations[i]
    }
    if amount > 0 {
        return []int{-1}
    }
    for i := range b {
        this.banknoteCounts[i] -= b[i]
    }
    return b
}


/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */