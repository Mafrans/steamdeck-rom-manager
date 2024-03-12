package me.mafrans.srm.ui

import androidx.compose.foundation.Image
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.interaction.collectIsHoveredAsState
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.size
import androidx.compose.material.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.graphics.toComposeImageBitmap
import androidx.compose.ui.res.loadImageBitmap
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import me.mafrans.srm.games.Game
import org.jetbrains.skia.Image
import kotlin.io.path.inputStream
import kotlin.io.path.readBytes

@Composable
fun LibraryItem(game: Game) {
    val interactionSource = remember { MutableInteractionSource() }
    val isHovered by interactionSource.collectIsHoveredAsState()

    TextButton(
            onClick = {},
            contentPadding = PaddingValues(0.dp),
            elevation = ButtonDefaults.elevation(if (isHovered) 8.dp else 0.dp)
    ) {
        LibraryItemCover(game.coverPath)
    }
}