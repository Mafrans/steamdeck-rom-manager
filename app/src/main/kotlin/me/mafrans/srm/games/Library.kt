package me.mafrans.srm.games

import com.akuleshov7.ktoml.file.TomlFileReader
import me.mafrans.srm.libraryDir
import java.io.File
import java.nio.file.Files
import java.nio.file.Path
import kotlin.io.path.exists
import kotlin.io.path.name
import kotlin.io.path.readBytes

class Library {
    var games: Set<Game> = hashSetOf()

    init {
        val toml = TomlFileReader()

        for (dir in Files.walk(libraryDir)) {
            val id = dir.name.toIntOrNull()
                    ?: continue

            val manifestPath = dir.resolve("manifest.toml")
            if (manifestPath.exists()) {
                games += toml.decodeFromFile(Game.serializer(), manifestPath.toString())
            }
        }

        println(games)
    }
}