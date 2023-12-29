package me.mafrans.srm.games

import GamesOuterClass.*
import java.io.ByteArrayOutputStream
import java.util.zip.Inflater

class GameDatabase(private val resource: String, private val isCompressed: Boolean = false) : HashMap<Int, Game>() {
    private val crcIndex = hashMapOf<Int, Int>()

    init {
        val games = readGames()
        for (game in games) {
            this[game.id] = game
            crcIndex[game.crcHash] = game.id
        }
    }

    fun getByCRC(crc: Int): Game? {
        val id = crcIndex[crc] ?: return null
        return this[id]
    }

    private fun readGames(): List<Game> {
        val stream = javaClass.getResourceAsStream(resource)
                ?: return emptyList()

        var bytes = stream.readAllBytes()
        if (isCompressed) {
            bytes = decompress(bytes)
        }

        val gamesContainer = Games.newBuilder().mergeFrom(bytes)
        return gamesContainer.gamesList
    }

    private fun decompress(input: ByteArray): ByteArray {
        val inflater = Inflater().apply { setInput(input) }
        val buf = ByteArray(1024)
        val outStream = ByteArrayOutputStream()
        outStream.use {
            do {
                val c = inflater.inflate(buf)
                it.write(buf, 0, c)
            } while (c != 0)
            inflater.end()
        }
        return outStream.toByteArray()
    }
}