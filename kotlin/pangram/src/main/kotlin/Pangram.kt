object Pangram {

    fun isPangram(input: String): Boolean {
        val allCharacters = ('a'..'z').toMutableSet()
        input.forEach { allCharacters.remove(it.toLowerCase()) }
        return allCharacters.isEmpty()

        // alt solution
        // val allCharacters = ('a'..'z').toSet()
        // val inputCharacters = input.toLowerCase().toSet()
        // return (allCharacters - inputCharacters).isEmpty()
    }
}
