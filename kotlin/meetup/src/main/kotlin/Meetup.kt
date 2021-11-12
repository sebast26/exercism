import java.time.DayOfWeek
import java.time.LocalDate
import java.time.YearMonth

class Meetup(month: Int, year: Int) {

    private val yearMonth = YearMonth.of(year, month)

    fun day(dayOfWeek: DayOfWeek, schedule: MeetupSchedule): LocalDate {
        var dayOffset = 1 + getWeekOffset(schedule) * 7

        val firstDay = yearMonth.atDay(1).dayOfWeek

        if (firstDay > dayOfWeek) {
            dayOffset += 7 - firstDay.value
        }


        val d = yearMonth.atDay(dayOffset)
        return d
    }

    private fun getWeekOffset(schedule: MeetupSchedule): Int {
        return when (schedule) {
            MeetupSchedule.FIRST -> 0
            MeetupSchedule.SECOND -> 1
            MeetupSchedule.THIRD -> 2
            MeetupSchedule.FOURTH -> 3
            MeetupSchedule.LAST -> TODO()
            MeetupSchedule.TEENTH -> TODO()
        }
    }
}
