type Bank struct {
    balances []int64
}


func Constructor(balance []int64) Bank {
    return Bank{balance}
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if !this.validateAccounts(account1, account2) || this.balances[account1 - 1] < money {
        return false
    }
    this.balances[account1 - 1] -= money
    this.balances[account2 - 1] += money
    return true
}


func (this *Bank) Deposit(account int, money int64) bool {
    if !this.validateAccounts(account) {
        return false
    }
    this.balances[account - 1] += money
    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if !this.validateAccounts(account) || this.balances[account - 1] < money {
        return false
    }
    this.balances[account - 1] -= money
    return true
}


func (this *Bank) validateAccounts(accounts ...int) bool {
    for _, account := range accounts {
        if account > len(this.balances) {
            return  false
        }
    }
    return true
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */