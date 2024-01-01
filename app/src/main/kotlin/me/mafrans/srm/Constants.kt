package me.mafrans.srm

import java.nio.file.Path
import kotlin.io.path.Path


const val appName = "steamdeck-rom-manager"
val homeDir = Path(System.getProperty("user.home"))
val appDir = determineAppDir()
val libraryDir = appDir.resolve("games")
val emulatorDir = appDir.resolve("emulators")

private fun determineAppDir(): Path {
    return when {
        OS.isUnix || OS.isMac -> homeDir.resolve(".local/share/$appName")
        OS.isWindows -> Path(System.getenv("APPDATA"), appName)
        else -> throw IllegalStateException("Unsupported operating system: ${OS.name}")
    }
}