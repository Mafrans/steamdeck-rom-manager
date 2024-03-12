package me.mafrans.srm.games

import com.akuleshov7.ktoml.Toml
import com.akuleshov7.ktoml.file.TomlFileWriter
import kotlinx.serialization.Serializable
import kotlinx.serialization.Transient
import me.mafrans.srm.libraryDir
import java.io.BufferedOutputStream
import java.io.FileOutputStream
import java.net.HttpURLConnection
import java.net.URI
import java.net.URL
import java.net.URLEncoder
import java.nio.charset.Charset
import kotlin.io.path.createDirectories
import kotlin.io.path.exists
import kotlin.io.path.outputStream

@Serializable
data class Game(
        var id: Int,
        var title: String,
        var console: Console,
) {
    @Transient
    val installPath = libraryDir.resolve(id.toString())

    @Transient
    val coverPath = installPath.resolve("cover.png")

    @Transient
    val romPath = installPath.resolve("binary.rom")

    @Transient
    val manifestPath = installPath.resolve("manifest.toml")

    @Transient
    var isInstalled = installPath.exists()

    fun install(bytes: ByteArray) {
        installPath.createDirectories()
        installBinary(bytes)
        writeManifest()

        isInstalled = true
        println("Installed '$title' into '$installPath'")

        // Download cover asynchronously
        Thread { downloadCover() }.start()
    }

    private fun installBinary(bytes: ByteArray) {
        val out = romPath.outputStream()
        out.write(bytes)
        out.close()
    }

    private fun writeManifest() {
        val writer = TomlFileWriter()
        writer.encodeToFile(serializer(), this, manifestPath.toString())
    }

    private fun downloadCover(title: String = this.title) {
        val encodedTitle = title.replace(" ", "%20")
        val url = "https://github.com/libretro-thumbnails/%s/raw/master/Named_Boxarts/%s.png"
                .format(console.title.replace(" ", "_"), encodedTitle)
                .let { URI(it).toURL() }

        with(url.openConnection() as HttpURLConnection) {
            val bytes = inputStream.readAllBytes()

            val contentString = bytes.toString(Charset.forName("UTF-8"))
            if (contentString.endsWith(".png")) {
                return downloadCover(contentString.removeSuffix(".png"))
            }

            val out = coverPath.outputStream()
            out.write(bytes)
            out.close()
        }
    }
}