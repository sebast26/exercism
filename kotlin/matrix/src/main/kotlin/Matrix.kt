class Matrix(private val matrixAsString: String) {

    private val whiteCharRegex = "\\s+".toRegex()
    private val rows = matrixAsString.splitToSequence(System.lineSeparator())

    fun column(colNr: Int): List<Int> {
        return rows
            .map { it.trim() }
            .map { it.split(whiteCharRegex) }
            .map { list -> list[colNr - 1] }
            .map { it.toInt() }
            .toList()
    }

    fun row(rowNr: Int): List<Int> {
        return rows
            .elementAt(rowNr - 1)
            .split(whiteCharRegex)
            .map { it.toInt() }
    }
}
