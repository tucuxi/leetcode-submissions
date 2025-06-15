
import java.util.Calendar

object Solution {
    def dayOfTheWeek(day: Int, month: Int, year: Int): String = {
        val cal = Calendar.getInstance
        cal.set(year, month - 1, day)
        cal.get(Calendar.DAY_OF_WEEK) match {
            case Calendar.SUNDAY => "Sunday"
            case Calendar.MONDAY => "Monday"
            case Calendar.TUESDAY => "Tuesday"
            case Calendar.WEDNESDAY => "Wednesday"
            case Calendar.THURSDAY => "Thursday"
            case Calendar.FRIDAY => "Friday"
            case Calendar.SATURDAY => "Saturday"
        }
    }
}