package me.mafrans.srm.ui

import androidx.compose.foundation.Image
import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.res.loadImageBitmap
import me.mafrans.srm.games.Game
import java.nio.file.Path
import kotlin.io.path.exists
import kotlin.io.path.inputStream

val defaultCover = loadImageBitmap(Game::class.java.getResourceAsStream("/default-cover.png")!!)

@Composable
fun LibraryItemCover(path: Path) {
    val bitmap = try {
        loadImageBitmap(path.inputStream())
    } catch (e: Exception) {
        defaultCover
    }


    Image(bitmap = bitmap, contentDescription = "")
}