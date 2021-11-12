import kotlin.experimental.and

class Allergies(number: Int) {

    private val score: Byte = number.toByte()

    fun getList() = Allergen.values().filter { scoreContainsAllergen(it) }

    fun isAllergicTo(allergen: Allergen) = scoreContainsAllergen(allergen)

    private fun scoreContainsAllergen(allergen: Allergen) = allergen.score.toByte().and(score) != 0.toByte()
}
