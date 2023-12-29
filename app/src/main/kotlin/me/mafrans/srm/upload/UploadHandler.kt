package me.mafrans.srm.upload

import io.javalin.http.Context
import io.javalin.http.Handler
import me.desair.tus.server.TusFileUploadService

class UploadHandler(val tus: TusFileUploadService) : Handler {
    override fun handle(ctx: Context) {
        tus.process(ctx.req(), ctx.res())
    }
}