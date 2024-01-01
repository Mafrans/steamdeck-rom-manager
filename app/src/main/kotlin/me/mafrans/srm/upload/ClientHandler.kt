package me.mafrans.srm.upload

import io.javalin.http.Context
import io.javalin.http.Handler
import io.javalin.http.NotFoundResponse
import java.nio.charset.Charset

class ClientHandler : Handler {
    var content: String?

    init {
        val stream = javaClass.getResourceAsStream("/index.html")
        content = stream?.readAllBytes()?.toString(Charset.forName("UTF-8"))
    }

    override fun handle(ctx: Context) {
        if (content != null) {
            ctx.html(content ?: "")
        } else throw NotFoundResponse()
    }
}