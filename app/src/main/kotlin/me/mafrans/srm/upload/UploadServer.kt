package me.mafrans.srm.upload

import io.javalin.Javalin
import io.javalin.apibuilder.ApiBuilder.*
import io.javalin.http.staticfiles.Location
import me.desair.tus.server.TusFileUploadService
import me.mafrans.srm.GAMES
import me.mafrans.srm.games.Game


class UploadServer() {
    val tus = TusFileUploadService()
            .withUploadUri("/api/upload")!!

    private val server: Javalin? = Javalin.create { config ->
        config.staticFiles.add { sf ->
            sf.hostedPath = "/assets"
            sf.directory = "/assets"
            sf.location = Location.CLASSPATH
        }
    }

    fun start() {
        val client = ClientHandler()
        val upload = UploadHandler().apply {
            onUpload = { handleUpload(it) }
        }

        if (server == null) {
            // TODO: Handle error message if server is null
            return
        }

        server.get("/", client).routes {
            path("/api/upload/") {
                post(upload)
                path("*") {
                    get(upload)
                    post(upload)
                    patch(upload)
                    head(upload)
                }
            }
        }.start()

        // Cleanup every 5 minutes
        Thread {
            while (true) {
                Thread.sleep(5 * 60 * 1000)
                cleanup()
            }
        }.start()
    }

    private fun handleUpload(upload: UploadedGame) {
        val identifiedGame = GAMES.getByCRC(upload.crcHash)
        println(identifiedGame?.name)
        if (identifiedGame == null) {
            // TODO: Handle unknown game
            return
        }

        val game = Game(identifiedGame.id, identifiedGame.name, identifiedGame.console)
        if (!game.isInstalled) {
            game.install(upload.bytes)
        }

        upload.close()
    }

    private fun cleanup() {
        tus.cleanup()
    }
}