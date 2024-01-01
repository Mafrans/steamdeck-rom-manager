package me.mafrans.srm.upload

import io.javalin.http.Context
import io.javalin.http.Handler
import me.mafrans.srm.SERVER


class UploadHandler() : Handler {
    var onUpload: (UploadedGame) -> Unit = {}
    override fun handle(ctx: Context) {
        val tus = SERVER.tus
        tus.process(ctx.req(), ctx.res())

        val uri = ctx.req().requestURI
        val upload = tus.getUploadInfo(uri)
        if (upload == null || upload.isUploadInProgress) {
            return
        }

        val byteStream = tus.getUploadedBytes(uri)
        val bytes = byteStream.readAllBytes()

        val game = UploadedGame(uri, upload, bytes)
        onUpload(game)

        tus.cleanup()
    }
}