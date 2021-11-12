object ScrabbleScore {

    fun scoreLetter(c: Char) = LETTER_SCORES.getOrDefault(c.toUpperCase(), 0)

    fun scoreWord(word: String) = word.sumBy { scoreLetter(it) }

    val LETTER_SCORES: Map<Char, Int> = mapOf(
        1 to listOf('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'),
        2 to listOf('D', 'G'),
        3 to listOf('B', 'C', 'M', 'P'),
        4 to listOf('F', 'H', 'V', 'W', 'Y'),
        5 to listOf('K'),
        8 to listOf('J', 'X'),
        10 to listOf('Q', 'Z')
    ).flatMap { e: Map.Entry<Int, List<Char>> -> e.value.associateWith { e.key }.toList() }.toMap()
}
