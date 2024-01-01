package me.mafrans.srm.upload

import com.google.common.math.IntMath.pow
import com.google.common.math.LongMath.log10
import me.desair.tus.server.upload.UploadInfo
import me.mafrans.srm.app
import java.util.zip.CRC32

private val crc = CRC32()

data class UploadedGame(val uri: String, val info: UploadInfo, val bytes: ByteArray) {
    val crcHash: Long

    init {
        crc.reset()
        crc.update(bytes)
        crcHash = crc.value
    }

    fun close() {
        app.server.tus.deleteUpload(uri)
    }

    override fun hashCode(): Int {
        var result = uri.hashCode()
        result = 31 * result + info.hashCode()
        return result
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as UploadedGame
        if (uri != other.uri) return false
        if (info != other.info) return false
        return true
    }

    override fun toString(): String {
        return """
            uri: $uri
            crcHash: $crcHash
            bytes: ${bytes.size}
        """.trimIndent()
    }
}