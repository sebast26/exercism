object Isogram {

    fun isIsogram(input: String): Boolean {
        val charactersFrequency = input
            .map { it.toLowerCase() }
            .filter { it in 'a'..'z' }
            .groupBy { it }
        return charactersFrequency.all { entry -> entry.value.size == 1 }
    }
}
