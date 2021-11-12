object Raindrops {

    fun convert(n: Int): String {
        val sb = StringBuilder()
        if (n % 3 == 0) {
            sb.append("Pling")
        }
        if (n % 5 == 0) {
            sb.append("Plang")
        }
        if (n % 7 == 0) {
            sb.append("Plong")
        }
        if (n % 3 != 0 && n % 5 != 0 && n % 7 != 0) {
            sb.append(n.toString())
        }
        return sb.toString()

        // alt 1
        // val factorAndWords = listOf(
        //     3 to "Pling",
        //     5 to "Plang",
        //     7 to "Plong"
        // )
        //
        // return factorAndWords
        //     .fold("", {acc, pair -> if (n.rem(pair.first) == 0) "$acc${pair.second}" else acc })
        //     .ifEmpty { n.toString() }
    }
}
