package me.mafrans.srm.games

import Games.*
import java.io.InputStream
import java.util.zip.GZIPInputStream

class GameDatabase(private val resource: String, private val isCompressed: Boolean = false) : HashMap<Int, GameDBEntry>() {
    private val crcIndex = hashMapOf<Long, Int>()

    init {
        val games = readGames()
        for (game in games) {
            this[game.id] = game
            crcIndex[game.crcHash] = game.id
        }

        val tetris = firstNotNullOf { (_, v) -> if (v.name.startsWith("Tetris (World)")) v else null }
    }

    fun getByCRC(crc: Long): GameDBEntry? {
        val id = crcIndex[crc] ?: return null
        return this[id]
    }

    private fun readGames(): List<GameDBEntry> {
        val stream = javaClass.getResourceAsStream(resource)
                ?: return emptyList()

        val bytes = if (isCompressed) decompress(stream) else stream.readAllBytes()

        val gamesContainer = GameDB.newBuilder().mergeFrom(bytes)
        return gamesContainer.gamesList
    }

    private fun decompress(input: InputStream): ByteArray {
        val gzip = GZIPInputStream(input)
        return gzip.readAllBytes()
    }
}