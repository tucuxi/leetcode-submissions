func coinChange(coins []int, amount int) int {
   dp := make([]int, amount + 1)
   for i := 1; i <= amount; i++ {
       dp[i] = amount + 1
       for _, c := range coins {
           if c <= i {
               if dp[i - c] < dp[i] {
                   dp[i] = dp[i - c] + 1
               }
           }
       }
   } 
   if dp[amount] > amount {
       return -1
   } 
   return dp[amount]
}