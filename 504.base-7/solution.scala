object Solution {
    def convertToBase7(num: Int): String = {
        def convAcc(acc: String, num: Int): String = {
            if (num == 0)
                acc
            else
                convAcc((num % 7).toString + acc, num / 7)
        }
        num match {
            case 0 => "0"
            case x if x < 0 => "-" + convAcc("", -x)
            case _ => convAcc("", num)
        }
    }
}