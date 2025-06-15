func rangeBitwiseAnd(left int, right int) int {
   res := left
   for i := left + 1; i <= right && res != 0; i++ {
       res &= i
   } 
   return res
}