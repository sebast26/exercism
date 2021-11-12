import kotlin.math.round

class SpaceAge(number: Int) {

    private val seconds = number.toDouble()

    enum class Planet(private val periodInEarthYears: Double) {
        EARTH(1.toDouble()),
        MERCURY(0.2408467),
        VENUS(0.61519726),
        MARS(1.8808158),
        JUPITER(11.862615),
        SATURN(29.447498),
        URANUS(84.016846),
        NEPTUNE(164.79132);

        fun countYears(seconds: Double): Double {
            return seconds.div(DAY_SECONDS * EARTH_YEAR * periodInEarthYears).round()
        }

        private fun Double.round() = round(this * 100) / 100

        companion object Constants {
            const val DAY_SECONDS = 60 * 60 * 24
            const val EARTH_YEAR = 365.25
        }
    }

    fun onEarth(): Double = Planet.EARTH.countYears(seconds)
    fun onMercury(): Double = Planet.MERCURY.countYears(seconds)
    fun onVenus(): Double = Planet.VENUS.countYears(seconds)
    fun onMars(): Double = Planet.MARS.countYears(seconds)
    fun onJupiter(): Double = Planet.JUPITER.countYears(seconds)
    fun onSaturn(): Double = Planet.SATURN.countYears(seconds)
    fun onUranus(): Double = Planet.URANUS.countYears(seconds)
    fun onNeptune(): Double = Planet.NEPTUNE.countYears(seconds)
}
