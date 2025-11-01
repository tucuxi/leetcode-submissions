class Bank(balance: LongArray) {

    private val balances = balance.copyOf()

    fun transfer(account1: Int, account2: Int, money: Long): Boolean {
        return (available(account1, money) && available(account2)).also {
            if (it) {
                balances[account1 - 1] -= money
                balances[account2 - 1] += money
            }
        }
    }

    fun deposit(account: Int, money: Long): Boolean {
        return available(account).also {
            if (it) {
                balances[account - 1] += money
            }
        }
    }

    fun withdraw(account: Int, money: Long): Boolean {
        return available(account, money).also {
            if (it) {
                balances[account - 1] -= money
            }
        }
    }

    private fun available(account: Int, amount: Long = 0): Boolean =
        account in 1 .. balances.size && balances[account - 1] >= amount
}

/**
 * Your Bank object will be instantiated and called as such:
 * var obj = Bank(balance)
 * var param_1 = obj.transfer(account1,account2,money)
 * var param_2 = obj.deposit(account,money)
 * var param_3 = obj.withdraw(account,money)
 */