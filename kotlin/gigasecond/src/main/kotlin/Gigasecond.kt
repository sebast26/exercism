import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(sourceDateTime: LocalDateTime) {

    constructor(sourceDate: LocalDate) : this(sourceDate.atStartOfDay())

    val date: LocalDateTime = sourceDateTime.plusSeconds(GIGASECONDS)

    companion object Constants {
        private val GIGASECONDS = 1E9.toLong()
    }
}
