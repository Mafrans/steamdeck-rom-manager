package me.mafrans.srm.upload

import io.javalin.Javalin
import io.javalin.apibuilder.ApiBuilder.*
import io.javalin.http.staticfiles.Location
import me.desair.tus.server.TusFileUploadService

class UploadServer {
    private val tus = TusFileUploadService()
    val server: Javalin? = Javalin.create { config ->
        config.staticFiles.add { sf ->
            sf.hostedPath = "/assets"
            sf.directory = "/assets"
            sf.location = Location.CLASSPATH
        }
    }

    fun start() {
        val upload = UploadHandler(tus)
        val client = ClientHandler()

        if (server == null) {
            // TODO: Handle error message if server is null
            return
        }

        server.get("/", client).routes {
            path("/api/upload") {
                get(upload)
                post(upload)
                patch(upload)
            }
        }.start()
    }
}